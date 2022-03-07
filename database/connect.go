package database

import (
	"fmt"

	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//connect to the db and iniatialize the db variable
	database, err := gorm.Open(mysql.Open("root:@/asset_be"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	fmt.Println("Connection Opened to Database")

	//Migrate the database
	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Asset{}, &models.Accesorie{}, &models.Department{}, &models.Admin{})
	fmt.Println("Database Migrated")
}
