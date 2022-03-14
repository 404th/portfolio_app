package ports

import (
	"github.com/404th/portfolio_app/models"
	"github.com/gofrs/uuid"
)

type APIPorts interface {
	GetSignUp(email, username string, password int32) (*models.User, error)
	GetSignIn(username string, password int32) (*models.User, error)
	GetDeleteUser(id uuid.UUID) (bool, error)
}
