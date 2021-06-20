package schema

import (
	"time"

	"github.com/jinzhu/gorm"
)

var (
	Database *gorm.DB
)

type TableInterface interface {
	Pk() string
	Ref() string
	AddForeignKeys()
	InsertDefaults()
}

type Base struct {
	Id        int        `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy string     `gorm:"default:null"`
	UpdatedBy string     `gorm:"default:null"`
	DeletedAt *time.Time `sql:"index"`
}

func AutoMigrate(database *gorm.DB) {
	Database = database

	database.AutoMigrate(
		&Links{},
		&Users{},
	)

	// Relationship

	// Insert Default

	// Update Table
}