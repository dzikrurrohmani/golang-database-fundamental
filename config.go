package main

import (
	"fmt"
	"go-db/utils"
	"os"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Db *sqlx.DB
}

func (c *Config) initDB() {
	dbHost := os.Getenv("TEMPDBHOST")
	dbPort := os.Getenv("TEMPDBPORT")
	dbName := os.Getenv("TEMPDBNAME")
	dbUser := os.Getenv("TEMPDBUSER")
	dbPassword := os.Getenv("TEMPDBPASSWORD")
	dbDriver := os.Getenv("TEMPDBDRIVER")

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s",dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect(dbDriver, dsn)
	utils.IsError(err)
	c.Db = db
}

func (c *Config) DbConn() *sqlx.DB {
	return c.Db
}

func NewConfigDB() Config {
	cfg := Config{}
	cfg.initDB()
	return cfg
}
