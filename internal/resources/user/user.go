package user

import (
	"context"
	"errors"
	"time"

	"github.com/sanpezlo/chess/internal/db"
	"go.uber.org/fx"
)

type Resource struct {
	dbClient *db.PrismaClient
}

func New(dbClient *db.PrismaClient) Repository {
	return &Resource{dbClient}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (r *Resource) CreateUser(ctx context.Context, email string, username string) (*User, error) {
	user, err := r.dbClient.User.CreateOne(
		db.User.Email.Set(email),
		db.User.Name.Set(username),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return FromModel(user, false), nil
}

func (r *Resource) GetUser(ctx context.Context, userId string, public bool) (*User, error) {
	user, err := r.dbClient.User.
		FindUnique(db.User.ID.Equals(userId)).
		Exec(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return FromModel(user, public), nil
}

func (r *Resource) GetUserByEmail(ctx context.Context, email string, public bool) (*User, error) {
	user, err := r.dbClient.User.
		FindUnique(db.User.Email.Equals(email)).
		Exec(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return FromModel(user, public), nil
}

func (r *Resource) GetUsers(ctx context.Context, sort string, max, skip int, public bool) ([]User, error) {
	users, err := r.dbClient.User.
		FindMany().
		Take(max).
		Skip(skip).
		OrderBy(
			db.User.CreatedAt.Order(db.Direction(sort)),
		).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return FromModelMany(users, public), nil
}

func (r *Resource) UpdateUser(ctx context.Context, userId string, email, name, bio *string) (*User, error) {
	user, err := r.dbClient.User.
		FindUnique(db.User.ID.Equals(userId)).
		Update(
			db.User.Email.SetIfPresent(email),
			db.User.Name.SetIfPresent(name),
			db.User.Bio.SetIfPresent(bio),
		).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return FromModel(user, false), nil
}

func (r *Resource) SetAdmin(ctx context.Context, userId string, status bool) (bool, error) {
	_, err := r.dbClient.User.FindUnique(db.User.ID.Equals(userId)).Update(
		db.User.Admin.Set(status),
	).Exec(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Resource) Ban(ctx context.Context, userId string) (*User, error) {
	user, err := r.dbClient.User.
		FindUnique(
			db.User.ID.Equals(userId),
		).
		Update(
			db.User.DeletedAt.Set(time.Now()),
		).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return FromModel(user, false), nil
}

func (r *Resource) Unban(ctx context.Context, userId string) (*User, error) {
	user, err := r.dbClient.User.
		FindUnique(
			db.User.ID.Equals(userId),
		).
		Update(
			db.User.DeletedAt.SetOptional(nil),
		).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return FromModel(user, false), nil
}
