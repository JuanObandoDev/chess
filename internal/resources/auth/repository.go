package auth

import (
	"context"
)

type Repository interface {
	CreateProvider(ctx context.Context, userID, accountID, username, email string, provider Provider) (*OAuthProvider, error)
	GetProvidersByUserId(ctx context.Context, userId string) ([]OAuthProvider, error)
}
