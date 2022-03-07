package database

import (
	"fmt"
	"log"

	"gitlab.ci.emalify.com/roamtech/asset_be/app/config"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//connect to the db and iniatialize the db variable
	database, err := gorm.Open(mysql.Open("root:@/asset_be"), &gorm.Config{})

	if err != nil {
		log.Println("cannot get port")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	//Migrate the database
	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Asset{}, &models.Accesorie{}, &models.Department{}, &models.Admin{})
	fmt.Println("Database Migrated")
}

func AutoMigrate() {
	err := DB.AutoMigrate(&models.Asset{}, &models.User{}, &models.Accesorie{}, &models.Department{})
	if err != nil {
		return
	}
	// fmt.Println("Database Migrated")

}
