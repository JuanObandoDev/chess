package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/sanpezlo/chess/internal/config"
	"github.com/sanpezlo/chess/internal/db"
	"github.com/sanpezlo/chess/internal/resources/auth"
	"github.com/sanpezlo/chess/internal/resources/user"
	"github.com/thanhpk/randstr"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type DiscordService struct {
	cache          *cache.Cache
	oa2Config      *oauth2.Config
	userRepository user.Repository
	authRepository auth.Repository
}

var _ AuthService = &DiscordService{}

func NewDiscordService(cfg *config.Config, userRepository user.Repository, authRepository auth.Repository) *DiscordService {
	var endpoint = oauth2.Endpoint{
		AuthURL:  "https://discord.com/api/oauth2/authorize",
		TokenURL: "https://discord.com/api/oauth2/token",
	}

	return &DiscordService{
		cache: cache.New(10*time.Minute, 20*time.Minute),
		oa2Config: &oauth2.Config{
			ClientID:     cfg.DiscordClientID,
			ClientSecret: cfg.DiscordClientSecret,
			Scopes:       []string{"identify", "email"},
			Endpoint:     endpoint,
		},
		userRepository: userRepository,
		authRepository: authRepository,
	}
}

func (gp *DiscordService) Link() string {
	state := randstr.String(16)
	err := gp.cache.Add(state, struct{}{}, 10*time.Minute)
	if err != nil {
		zap.L().Error("failed to add state to cache", zap.Error(err))
	}
	return gp.oa2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (gp *DiscordService) Login(ctx context.Context, state, code string) (*user.User, error) {

	if _, ok := gp.cache.Get(state); !ok {
		return nil, ErrStateMismatch
	}

	token, err := gp.oa2Config.Exchange(ctx, code)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform OAuth2 token exchange")
	}

	//

	req, err := http.NewRequest("GET", "https://discordapp.com/api/users/@me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("discord responded with %d", resp.StatusCode)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var discordUser discordgo.User
	err = json.NewDecoder(bytes.NewReader(content)).Decode(&discordUser)
	if err != nil {
		return nil, err
	}

	//

	email := discordUser.Email
	if email == "" {
		return nil, errors.New("email missing from Discord account data")
	}

	u, err := gp.userRepository.GetUserByEmail(ctx, email, false)
	if err != nil && err != db.ErrNotFound {
		return nil, err
	}

	if u == nil {
		u, err = gp.userRepository.CreateUser(ctx, email, discordUser.Username, discordUser.AvatarURL(""))
		if err != nil {
			return nil, err
		}
	}
	if findOauthProvider(u, auth.Discord) == nil {
		if _, err := gp.authRepository.CreateProvider(ctx, u.ID, fmt.Sprint(discordUser.ID), discordUser.Username, discordUser.Email, auth.Discord); err != nil {
			return nil, errors.Wrap(err, "failed to create user Discord relationship")
		}

	}

	return u, nil
}
