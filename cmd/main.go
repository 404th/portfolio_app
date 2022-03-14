package main

import (
	"log"

	config "github.com/404th/portfolio_app/Config"
	"github.com/404th/portfolio_app/internal/adapters/app/api"
	"github.com/404th/portfolio_app/internal/adapters/framework/left/grpc"
	"github.com/404th/portfolio_app/internal/adapters/framework/right/db"
	_ "github.com/lib/pq"
)

func main() {
	server_port := ":6464"
	cfg := &config.DBCfg{
		PgDriverName: "postgres",
		PgHost:       "localhost",
		PgSSLMode:    "disable",
		PgPort:       "5432",
		PgDBName:     "registered_users",
		PgUser:       "postgres",
		PgPassword:   "postgres",
	}

	pg_db, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		return
	}
	defer pg_db.Close()

	// adapter
	user_adapter := api.NewAdapter(pg_db)
	new_serv := grpc.NewAdapter(user_adapter)

	if err = new_serv.Run(server_port); err != nil {
		log.Fatalf("Failed to run server over port%v and occured this error:%v", server_port, err)
		return
	}
}
