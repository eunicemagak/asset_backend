package database

import (
	"fmt"
	"log"
	"strconv"

	"gitlab.ci.emalify.com/roamtech/asset_be/app/config"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

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

	// Migrate the database
	DB.AutoMigrate(&models.Asset{}, &models.User{}, &models.Assign{}, &models.Accesorie{}, &models.Department{})
	fmt.Println("Database Migrated")
}

func AutoMigrate() {
	err := DB.AutoMigrate(&models.Asset{}, &models.User{}, &models.Accesorie{}, &models.Department{})
	if err != nil {
		return
	}
	// fmt.Println("Database Migrated")

}
