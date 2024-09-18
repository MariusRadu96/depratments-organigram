package config

import (
	"sync"

	"github.com/jinzhu/configor"
)

type Config struct {
	Port  int    `env:"PORT" default:"8081"`
	Host  string `env:"HOST" default:"0.0.0.0"`
	IsDev bool   `env:"IS_DEV" default:"false"`

	// DB
	MySQLDB       string `env:"MYSQL_DATABASE" required:"true"`
	MySQLPort     int    `env:"MYSQL_PORT" default:"5432"`
	MySQLHost     string `env:"MYSQL_HOST" required:"true"`
	MySQLUser     string `env:"MYSQL_USER" required:"true"`
	MySQLPassword string `env:"MYSQL_PASSWORD" required:"true"`

	JWTSecret string `env:"JWT_SECRET" required:"true"`
}

func initConfig() {
	if err := configor.New(&configor.Config{}).Load(&conf, ""); err != nil {
		panic(err)
	}

}

var (
	conf Config
	once sync.Once
)

func Get() Config {
	once.Do(initConfig)
	return conf
}
