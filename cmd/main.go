package main

import (
	"log"
	"os"

	config "github.com/404th/portfolio_app/Config"
	"github.com/404th/portfolio_app/internal/adapters/app/api"
	"github.com/404th/portfolio_app/internal/adapters/framework/left/grpc"
	client "github.com/404th/portfolio_app/internal/adapters/framework/left/grpc/grpc_client"
	"github.com/404th/portfolio_app/internal/adapters/framework/right/db"
	_ "github.com/lib/pq"
)

func main() {
	os.Setenv("PG_DRIVER_NAME", "postgres")
	os.Setenv("PG_HOST", "localhost")
	os.Setenv("PG_SSL_MODE", "disable")
	os.Setenv("PG_PORT", "5432")
	os.Setenv("PG_DB_NAME", "registered_users")
	os.Setenv("PG_USER", "postgres")
	os.Setenv("PG_PASSWORD", "postgres")
	os.Setenv("SERVER_PORT", "0.0.0.0:6464")

	Cfg := &config.DBCfg{
		PgDriverName: os.Getenv("PG_DRIVER_NAME"),
		PgHost:       os.Getenv("PG_HOST"),
		PgSSLMode:    os.Getenv("PG_SSL_MODE"),
		PgPort:       os.Getenv("PG_PORT"),
		PgDBName:     os.Getenv("PG_DB_NAME"),
		PgUser:       os.Getenv("PG_USER"),
		PgPassword:   os.Getenv("PG_PASSWORD"),
	}

	client, err := client.New()
	if err != nil {
		log.Fatalf("Failed to get client: %s", err.Error())
		return
	}
	// orderService := service.NewOrderService(connDb, log, client)

	pg_db, err := db.NewDB(Cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		return
	}
	defer pg_db.Close()

	// adapter
	user_adapter := api.NewAdapter(pg_db)
	new_serv := grpc.NewAdapter(user_adapter)

	if err = new_serv.Run(os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalf("Failed to run server over port%v and occured this error:%v", os.Getenv("SERVER_PORT"), err)
		return
	}
}
