package auth

import (
	"github.com/sanpezlo/chess/internal/db"
)

type OAuthProvider struct {
	UserId    string `json:"userId"`
	User      string `json:"user"`
	AccountId string `json:"accountId"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Provider  string `json:"provider"`
}

func FromModel(o *db.OAuthProviderModel) *OAuthProvider {

	oauth := &OAuthProvider{
		UserId:    o.InnerOAuthProvider.UserID,
		User:      o.RelationsOAuthProvider.User.ID,
		AccountId: o.InnerOAuthProvider.AccountID,
		Username:  o.InnerOAuthProvider.Username,
		Email:     o.InnerOAuthProvider.Email,
		Provider:  string(o.InnerOAuthProvider.Provider),
	}

	return oauth
}

func FromModelMany(u []db.OAuthProviderModel) []OAuthProvider {
	users := make([]OAuthProvider, 0, len(u))
	for _, v := range u {
		users = append(users, *FromModel(&v))
	}
	return users
}
