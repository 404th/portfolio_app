package api

import (
	"github.com/404th/portfolio_app/internal/ports"
	"github.com/gofrs/uuid"
)

type Adapter struct {
	db   ports.DBPort
	user ports.UserPorts
}

func NewAdapter(db ports.DBPort, user ports.UserPorts) *Adapter {
	return &Adapter{db, user}
}

func (apia *Adapter) GetSignUp(email, username string, password int32) (uuid.UUID, error) {
	var signed_in_user_id uuid.UUID

	n_u, err := apia.user.SignIn(username, password)
	if err != nil {
		return signed_in_user_id, err
	}

	signed_in_user_id = n_u
	return signed_in_user_id, nil
}

func (apia *Adapter) GetSignIn(username string, password int32) (uuid.UUID, error) {
	var new_user_id uuid.UUID

	n_u, err := apia.user.SignIn(username, password)
	if err != nil {
		return new_user_id, err
	}

	new_user_id = n_u
	return new_user_id, nil
}

func (apia *Adapter) GetDeleteUser(id uuid.UUID) (bool, error) {
	is_deleted, err := apia.user.DeleteUser(id)
	if err != nil {
		return is_deleted, err

	}
	return is_deleted, nil
}
