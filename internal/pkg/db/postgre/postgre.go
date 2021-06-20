package database

import (
	"database/sql"
	"fmt"
	"github.com/glyphack/graphlq-golang/config"
	"github.com/glyphack/graphlq-golang/internal/pkg/db/schema"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var (
	Client *sql.DB
	Database *gorm.DB
)

func InitDB() {
	db, err := gorm.Open(
		config.ENV.Db.Driver,
		fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			config.ENV.Db.Host,
			config.ENV.Db.User,
			config.ENV.Db.DbName,
			config.ENV.Db.Password),
	)
	if err != nil {
		log.Panic(err)
	}
	db.LogMode(true)

	Client = db.DB()
	Database = db
}

func Migrate() {
	if config.ENV.Migration.DbAutoCreate == true {
		fmt.Println("Dropping and recreating all tables...")
		schema.AutoMigrate(Database)
		fmt.Println("All tables recreated successfully...")
	}

}
