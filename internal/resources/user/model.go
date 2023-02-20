package user

import (
	"time"

	"github.com/sanpezlo/chess/internal/db"
)

type User struct {
	ID    string  `json:"id"`
	Email string  `json:"email"`
	Name  string  `json:"name"`
	Bio   *string `json:"bio"`
	Admin bool    `json:"admin"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func FromModel(u *db.UserModel, public bool) *User {
	user := &User{
		ID:    u.InnerUser.ID,
		Email: u.InnerUser.Email,

		Name:      u.InnerUser.Name,
		Bio:       u.InnerUser.Bio,
		Admin:     u.InnerUser.Admin,
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
