package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/thephiri/go-vue/models"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:WordPass@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}
