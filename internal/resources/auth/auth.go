package auth

import (
	"context"
	"errors"

	"github.com/sanpezlo/chess/internal/db"
	"go.uber.org/fx"
)

type Provider int

const (
	GitHub Provider = iota
	Discord
)

func (p Provider) String() string {
	switch p {
	case GitHub:
		return "GITHUB"
	case Discord:
		return "DISCORD"
	default:
		return "UNDEFINED"
	}
}

type Resource struct {
	dbClient *db.PrismaClient
}

func New(dbClient *db.PrismaClient) Repository {
	return &Resource{dbClient}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (r *Resource) CreateProvider(ctx context.Context, userID, accountID, username, email string, provider Provider) (*OAuthProvider, error) {

	aouthProvider, err := getAuthProvider(provider)
	if err != nil {
		return nil, err
	}

	o, err := r.dbClient.OAuthProvider.CreateOne(
		db.OAuthProvider.User.Link(db.User.ID.Equals(userID)),
		db.OAuthProvider.AccountID.Set(accountID),
		db.OAuthProvider.Username.Set(username),
		db.OAuthProvider.Email.Set(email),
		db.OAuthProvider.Provider.Set(aouthProvider),
	).With(
		db.OAuthProvider.User.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return FromModel(o), nil

}

func (r *Resource) GetProvidersByUserId(ctx context.Context, userID string) ([]OAuthProvider, error) {

	o, err := r.dbClient.OAuthProvider.FindMany(
		db.OAuthProvider.User.Link(db.User.ID.Equals(userID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return FromModelMany(o), nil

}

func getAuthProvider(provider Provider) (db.AuthProvider, error) {
	switch provider {
	case GitHub:
		return db.AuthProviderGITHUB, nil
	case Discord:
		return db.AuthProviderDISCORD, nil
	default:
		return db.AuthProviderUNDEFINED, errors.New("invalid provider")
	}
}
