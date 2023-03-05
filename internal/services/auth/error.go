package auth

import "github.com/pkg/errors"

var (
	ErrStateMismatch = errors.New("state nonce mismatch")
	ErrOAuthNoEmail  = errors.New("missing email address on OAuth provider account")
)
