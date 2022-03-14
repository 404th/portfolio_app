package ports

import (
	"github.com/404th/portfolio_app/models"
	"github.com/gofrs/uuid"
)

type UserPorts interface {
	SignUp(email, username string, password int32) (*models.User, error)
	SignIn(username string, password int32) (*models.User, error)
	DeleteUser(id uuid.UUID) (bool, error)
	UpdateUser(new_username string, new_password int32, id uuid.UUID) (*models.User, error)
}
