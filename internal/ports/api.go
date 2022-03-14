package ports

import (
	"github.com/gofrs/uuid"
)

type APIPorts interface {
	GetSignUp(email, username string, password int32) (uuid.UUID, error)
	GetSignIn(username string, password int32) (uuid.UUID, error)
	GetDeleteUser(id uuid.UUID) (bool, error)
}
