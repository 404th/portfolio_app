package config

type DBCfg struct {
	PgDriverName string
	PgPort       string
	PgHost       string
	PgSSLMode    string
	PgDBName     string
	PgPassword   string
	PgUser       string
}

func NewDBCfg(driver_name, port, host, ssl_mode, db_name, password string) *DBCfg {
	return &DBCfg{
		PgDriverName: driver_name,
		PgPort:       port,
		PgHost:       host,
		PgSSLMode:    ssl_mode,
		PgDBName:     db_name,
		PgPassword:   password,
	}
}
