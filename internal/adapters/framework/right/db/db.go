package db

import (
	"fmt"

	config "github.com/404th/portfolio_app/Config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sqlx.DB
}

func NewDB(cfg *config.DBCfg) (*DB, error) {
	// making connection string for PostgreSQL
	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.PgHost,
		cfg.PgPort,
		cfg.PgUser,
		cfg.PgPassword,
		cfg.PgDBName,
		cfg.PgSSLMode,
	)
	new_db, err := sqlx.Connect(cfg.PgDriverName, conStr)
	if err != nil {
		return nil, err
	}

	if err = new_db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db: new_db}, nil
}

func (nd *DB) CloseConnection() error {
	return nd.db.Close()
}
