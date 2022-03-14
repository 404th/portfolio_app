package ports

import (
	"github.com/google/uuid"
)

type APIPorts interface {
	GetSignUp(email, username string, password int32) (uuid.UUID, error)
	GetSignIn(username string, password int32) (string, string, error)
	GetDeleteUser(id uuid.UUID) (bool, error)
}
