package conn

import (
	"go-fiber/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/gin"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Barang{})

	DB = database
}