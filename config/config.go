package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
)

type configuration struct {
	Debug      bool   `json:"debug"`
	Port       string `json:"port"`
	AppDebug   bool   `json:"app_debug"`
	Pool       *cache.Cache

	Db struct {
		Driver string `json:"driver"`
		User string `json:"user"`
		Password string `json:"password"`
		Host string `json:"host"`
		Port     string    `json:"port"`
		DbName       string `json:"db_name"`
	} `json:"db"`

	Migration struct{
		DbAutoCreate bool `json:"db_auto_create"`
	}

	Auth struct {
		SecretKey        string        `json:"secret_key"`
	} `json:"auth"`
}

var ENV configuration

func LoadENV() error {
	_ = godotenv.Load()
	if err := mapToENV(); err != nil {
		return err
	}
	return nil
}

func mapToENV() (err error) {
	// general
	ENV.AppDebug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	ENV.Pool = cache.New(6*time.Hour, 10*time.Minute)
	ENV.Port = os.Getenv("PORT")

	// postgre
	ENV.Db.Driver = os.Getenv("DB_DRIVER")
	ENV.Db.User = os.Getenv("DB_USER")
	ENV.Db.Password = os.Getenv("DB_PASS")
	ENV.Db.Host = os.Getenv("DB_HOST")
	ENV.Db.Port = os.Getenv("DB_PORT")
	ENV.Db.DbName = os.Getenv("DB_NAME")

	// migration
	ENV.Migration.DbAutoCreate, _ = strconv.ParseBool(os.Getenv("DB_AUTO_CREATE"))

	return
}