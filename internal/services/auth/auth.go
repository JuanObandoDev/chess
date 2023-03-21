package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/sanpezlo/chess/internal/config"
	"github.com/sanpezlo/chess/internal/resources/auth"
	"github.com/sanpezlo/chess/internal/resources/user"
	"github.com/sanpezlo/chess/internal/web"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const secureCookieName = "x-session"

type Cookie struct {
	UserID  string
	Admin   bool
	Created time.Time
}

type Service struct {
	sc *securecookie.SecureCookie
}

func NewService(cfg *config.Config) *Service {
	return &Service{sc: securecookie.New(cfg.HashKey, cfg.BlockKey)}
}

func (s *Service) EncodeAuthCookie(w http.ResponseWriter, user *user.User) {
	encoded, err := s.sc.Encode(secureCookieName, Cookie{
		UserID:  user.ID,
		Admin:   user.Admin,
		Created: time.Now(),
	})
	if err != nil {
		web.StatusUnauthorized(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     secureCookieName,
		Value:    encoded,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	})
}

func (s *Service) DecodeAuthCookie(r *http.Request, auth *Auth) bool {
	cookie, err := r.Cookie(secureCookieName)
	if err != nil {
		return false
	}

	if err = s.sc.Decode(secureCookieName, cookie.Value, auth.Cookie); err != nil {
		zap.L().Debug("failed to decode auth cookie", zap.Error(err))
		return false
	}

	return true
}

var Module = fx.Options(
	fx.Provide(NewService),
	fx.Provide(NewGitHubService),
	fx.Provide(NewDiscordService),
)

type AuthService interface {
	Link() string
	Login(ctx context.Context, state, code string) (*user.User, error)
}

func findOauthProvider(u *user.User, provider auth.Provider) *user.OAuthProvider {
	for _, p := range u.OAuthProviders {
		if p.Provider == provider.String() {
			return &p
		}
	}

	return nil
}
