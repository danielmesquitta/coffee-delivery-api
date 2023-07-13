package controller

import (
	"github.com/danielmesquitta/coffee-delivery-api/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = config.GetDatabase()
}
