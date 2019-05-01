package migration

import (
	"eduSense/configs"
	"eduSense/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Migrate() {
	db, err := gorm.Open(configs.DB_DIALECT, configs.DB_NAME)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	fmt.Printf("Migration Starting")

	// Migrate the schema
	db.AutoMigrate(&models.Student{})

	fmt.Printf("Migration Done")
}
