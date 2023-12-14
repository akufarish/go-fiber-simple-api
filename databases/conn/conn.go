package conn

import (
	"go-fiber/databases/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/gin"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&migrations.User{})

	DB = database
}