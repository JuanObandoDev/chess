package user

import (
	"time"

	"github.com/sanpezlo/chess/internal/db"
)

type OAuthProvider struct {
	Username string `json:"username"`
	Provider string `json:"provider"`
}

type User struct {
	ID             string          `json:"id"`
	Email          string          `json:"email"`
	Name           string          `json:"name"`
	Avatar         string          `json:"avatar"`
	Bio            *string         `json:"bio"`
	Admin          bool            `json:"admin"`
	OAuthProviders []OAuthProvider `json:"oauthProvider"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func FromModel(u *db.UserModel, public bool) *User {
	var oauthProviders []OAuthProvider

	if u.RelationsUser.OauthProviders != nil {
		oauthProviders = make([]OAuthProvider, 0, len(u.RelationsUser.OauthProviders))
		for _, v := range u.RelationsUser.OauthProviders {
			oauthProviders = append(oauthProviders, OAuthProvider{
				Username: v.Username,
				Provider: string(v.Provider),
			})
		}
	}

	user := &User{
		ID:             u.InnerUser.ID,
		Email:          u.InnerUser.Email,
		Name:           u.InnerUser.Name,
		Avatar:         u.InnerUser.Avatar,
		Bio:            u.InnerUser.Bio,
		Admin:          u.InnerUser.Admin,
		OAuthProviders: oauthProviders,

		CreatedAt: u.InnerUser.CreatedAt,
		UpdatedAt: u.InnerUser.UpdatedAt,
		DeletedAt: u.InnerUser.DeletedAt,
	}

	if public {
		user.Email = ""
	}

	return user
}

func FromModelMany(u []db.UserModel, public bool) []User {
	users := make([]User, 0, len(u))
	for _, v := range u {
		users = append(users, *FromModel(&v, public))
	}
	return users
}
