package main

import (
	"log"

	config "github.com/404th/portfolio_app/Config"
	"github.com/404th/portfolio_app/internal/adapters/framework/right/db"
	_ "github.com/lib/pq"
)

func main() {

	cfg := &config.DBCfg{
		PgDriverName: "postgres",
		PgHost:       "localhost",
		PgSSLMode:    "verify-full",
		PgPort:       "5432",
		PgDBName:     "registered_users",
		PgUser:       "postgres",
		PgPassword:   "password",
	}

	pg_db, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		return
	}

}
