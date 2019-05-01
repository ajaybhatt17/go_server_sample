package helpers

import (
	"eduSense/configs"
	"github.com/jinzhu/gorm"
)


func GetDBConnection() (db *gorm.DB) {
	db, err := gorm.Open(configs.DB_DIALECT, configs.DB_NAME)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
