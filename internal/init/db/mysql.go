package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string
	Name     string
	Port     int
	User     string
	Password string
}

func NewMysqlDB(config *Config) *sql.DB {
	formattedURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := sql.Open("mysql", formattedURL)
	if err != nil {
		log.Fatalf("Unable to connect to mysql: %v", err)
	}

	return db
}
