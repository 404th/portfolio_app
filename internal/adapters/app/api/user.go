package api

import (
	"errors"

	"github.com/404th/portfolio_app/internal/ports"
	"github.com/404th/portfolio_app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserAdapter struct {
	db  *sqlx.DB
	api ports.APIPorts
}

func NewAdapter(db *sqlx.DB, api ports.APIPorts) *UserAdapter {
	return &UserAdapter{db, api}
}

func (ua *UserAdapter) GetSignUp(email, username string, password int32) (uuid.UUID, error) {
	var null_uuid uuid.UUID
	qs := `
		INSERT INTO registered_users
			(
				uuid,
				email,
				username,
				password
			) VALUES (
				$1, $2, $3, $4
			)
	`
	// generating UUID
	uuid := uuid.New()
	if _, err := ua.db.Exec(qs, uuid, email, username, password); err != nil {
		return null_uuid, err
	}

	return uuid, nil
}

func (ua *UserAdapter) GetSignIn(username string, password int32) (string, string, error) {
	var signed_in_user *models.User
	find_user := `
		SELECT username, password FROM registered_users WHERE username=$1
	`
	row := ua.db.QueryRow(find_user, username)
	if err := row.Scan(&signed_in_user.Username, &signed_in_user.Password); err != nil {
		return "", "", err
	}

	// checking password
	if signed_in_user.Password != password {
		return "", "", errors.New("password is not compatible")
	}

	return signed_in_user.Username, signed_in_user.Email, nil
}

func (ua *UserAdapter) GetDeleteUser(id uuid.UUID) (bool, error) {
	qs := `
		DELETE FROM registered_users WHERE uuid=$1
	`

	_, err := ua.db.Exec(qs, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
