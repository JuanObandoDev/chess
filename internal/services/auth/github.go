package auth

import (
	"context"
	"fmt"
	"time"

	github_api "github.com/google/go-github/v50/github"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/sanpezlo/chess/internal/config"
	"github.com/sanpezlo/chess/internal/db"
	"github.com/sanpezlo/chess/internal/resources/auth"
	"github.com/sanpezlo/chess/internal/resources/user"
	"github.com/thanhpk/randstr"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GitHubService struct {
	cache          *cache.Cache
	oa2Config      *oauth2.Config
	userRepository user.Repository
	authRepository auth.Repository
}

var _ AuthService = &GitHubService{}

func NewGitHubService(cfg *config.Config, userRepository user.Repository, authRepository auth.Repository) *GitHubService {
	return &GitHubService{
		cache: cache.New(10*time.Minute, 20*time.Minute),
		oa2Config: &oauth2.Config{
			ClientID:     cfg.GithubClientID,
			ClientSecret: cfg.GithubClientSecret,
			Scopes:       []string{"read:user", "user:email"},
			Endpoint:     github.Endpoint,
		},
		userRepository: userRepository,
		authRepository: authRepository,
	}
}

func (gp *GitHubService) Link() string {
	state := randstr.String(16)
	err := gp.cache.Add(state, struct{}{}, 10*time.Minute)
	if err != nil {
		zap.L().Error("failed to add state to cache", zap.Error(err))
	}
	return gp.oa2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (gp *GitHubService) Login(ctx context.Context, state, code string) (*user.User, error) {

	if _, ok := gp.cache.Get(state); !ok {
		return nil, ErrStateMismatch
	}

	token, err := gp.oa2Config.Exchange(ctx, code)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform OAuth2 token exchange")
	}

	client := github_api.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token.AccessToken})))
	githubUser, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get GitHub user data")
	}

	email := githubUser.GetEmail()
	if email == "" {
		return nil, errors.New("email missing from GitHub account data")
	}

	u, err := gp.userRepository.GetUserByEmail(ctx, email, false)
	if err != nil && err != db.ErrNotFound {
		return nil, err
	}

	if u == nil {
		u, err = gp.userRepository.CreateUser(ctx, email, githubUser.GetLogin())
		if err != nil {
			return nil, err
		}
	}
	if findOauthProvider(u, auth.GitHub) == nil {
		if _, err := gp.authRepository.CreateProvider(ctx, u.ID, fmt.Sprint(githubUser.GetID()), githubUser.GetLogin(), githubUser.GetEmail(), auth.GitHub); err != nil {
			return nil, errors.Wrap(err, "failed to create user GitHub relationship")
		}

	}

	return u, nil
}
