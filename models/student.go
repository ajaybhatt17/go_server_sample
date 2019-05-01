package models


import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
}
