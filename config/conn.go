package config

import (
	"enigmacamp.com/go-db-fundamnetal/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Config struct {
	Db *sqlx.DB
}

func (c *Config) initDb() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")
	// urutan url koneksi ke db postgres
	// postgres://dbUser:dbPassword@dbHost:dbPort/dbName
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect(dbDriver, dsn)
	utils.IsError(err)
	c.Db = db
}

func (c *Config) DbConn() *sqlx.DB {
	return c.Db
}

func NewConfigDB() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}
