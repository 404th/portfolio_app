package ports

import (
	"github.com/gofrs/uuid"
)

type UserPorts interface {
	SignUp(email, username string, password int32) (uuid.UUID, error)
	SignIn(username string, password int32) (uuid.UUID, error)
	DeleteUser(id uuid.UUID) (bool, error)
}
