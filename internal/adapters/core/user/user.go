package user

import (
	"log"

	"github.com/404th/portfolio_app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Adapter struct {
	db *sqlx.DB
}

func NewAdapter(db *sqlx.DB) *Adapter {
	return &Adapter{db}
}

func (ua Adapter) SignUp(email, username string, password int32) (*models.User, error) {
	var new_user *models.User
	//
	new_id, err := uuidGen()
	if err != nil {
		return new_user, err
	}

	var user *models.User = &models.User{
		UUID:     new_id,
		Username: username,
		Email:    email,
		Password: password,
	}

	ins_str := `
		INSERT INTO registered_users (
			uuid,
			username,
			email,
			password,
			deleted
		) VALUES (
			$1, $2, $3, $4, $5,
		)
	`
	_, err = ua.db.Exec(ins_str, user.UUID, user.Username, user.Email, user.Password, false)
	if err != nil {
		return new_user, err
	}
	new_user = user
	return new_user, nil
}

func (ua Adapter) SignIn(username string, password int32) (*models.User, error) {
	var new_user *models.User

	find_str := `SELECT email, uuid, deleted FROM registered_users WHERE username=$1 AND password=$2`
	row := ua.db.QueryRow(find_str)

	if err := row.Scan(
		&new_user.Email,
		&new_user.UUID,
	); err != nil {
		return new_user, err
	}

	return new_user, nil
}

func (ua Adapter) DeleteUser(id uuid.UUID) (bool, error) {
	qs := `UPDATE registered_users SET deleted=$1 WHERE uuid=$2`

	_, err := ua.db.Exec(qs, true, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func uuidGen() (uuid.UUID, error) {
	nuuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Internal error while creating UUID for user: %v", err)
	}

	return nuuid, nil
}
