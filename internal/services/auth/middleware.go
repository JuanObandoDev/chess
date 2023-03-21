package auth

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sanpezlo/chess/internal/web"
)

type Auth struct {
	Authenticated bool
	Cookie        *Cookie
}

var contextKey = struct{}{}

func (s *Service) WithAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := &Auth{
			Authenticated: false,
			Cookie:        &Cookie{},
		}

		if s.DecodeAuthCookie(r, auth) {
			auth.Authenticated = true
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(
			r.Context(),
			contextKey,
			auth,
		)))
	})
}

func MustBeAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, ok := GetAuthenticationInfo(w, r)
		if !ok {
			return
		}
		if !auth.Authenticated {
			web.StatusUnauthorized(w, web.WithSuggestion(
				errors.New("user not authenticated"),
				"The request did not have any authentication information with it.",
				"Ensure you are logged in, try logging out and back in again. If issues persist, please contact us.",
			))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func MustBeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, ok := GetAuthenticationInfo(w, r)
		if !ok {
			return
		}

		if !auth.Cookie.Admin {
			web.StatusUnauthorized(w, errors.New("user is not an administrator"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetAuthenticationInfo(
	w http.ResponseWriter,
	r *http.Request,
) (*Auth, bool) {

	if auth, ok := r.Context().Value(contextKey).(*Auth); ok {
		return auth, true
	}

	web.StatusInternalServerError(w, web.WithSuggestion(
		errors.New("failed to extract auth context from request"),
		"Could not read session data from cookies.",
		"Try clearing your cookies and logging in to your account again."))
	return nil, false
}
