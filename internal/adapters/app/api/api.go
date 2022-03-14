package api

import (
	"github.com/404th/portfolio_app/internal/adapters/framework/right/db"
	"github.com/404th/portfolio_app/internal/ports"
	"github.com/404th/portfolio_app/models"
	"github.com/gofrs/uuid"
)

type Adapter struct {
	db   *db.DB
	user ports.UserPorts
}

func NewAdapter(db *db.DB, user ports.UserPorts) *Adapter {
	return &Adapter{db, user}
}

func (apia *Adapter) GetSignUp(email, username string, password int32) (*models.User, error) {
	var signed_in_user *models.User

	n_u, err := apia.user.SignIn(username, password)
	if err != nil {
		return signed_in_user, err
	}

	signed_in_user = n_u
	return signed_in_user, nil
}

func (apia *Adapter) GetSignIn(username string, password int32) (*models.User, error) {
	var new_user *models.User

	n_u, err := apia.user.SignUp(email, username, password)
	if err != nil {
		return new_user, err
	}

	new_user = n_u
	return new_user, nil
}

func (apia *Adapter) GetDeleteUser(id uuid.UUID) (bool, error) {
	is_deleted, err := apia.user.DeleteUser(id)
	if err != nil {
		return is_deleted, err

	}
	return is_deleted, nil
}
