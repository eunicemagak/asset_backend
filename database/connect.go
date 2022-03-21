package database

//USING SQL

// import (
// 	"fmt"

// 	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connect() {
// 	//connect to the db and iniatialize the db variable
// 	database, err := gorm.Open(mysql.Open("root:@/asset_be"), &gorm.Config{})

// 	if err != nil {
// 		panic("Could not connect to the database")
// 	}
// 	fmt.Println("Connection Opened to Database")

// 	//Migrate the database
// 	DB = database
// 	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Accesorie{}, &models.Department{}, &models.Admin{})
// 	fmt.Println("Database Migrated")
// }

//USING POSTGRES
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
<<<<<<< HEAD

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	AutoMigrate()

}

// Migrate the database

func AutoMigrate() {
	err := DB.AutoMigrate(&models.Admin{}, &models.User{}, &models.Department{}, &models.Asset{}, &models.Accesorie{}, &models.Image{})
=======

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&models.Asset{}, &models.User{}, &models.Role{}, &models.Accesorie{}, &models.Department{}, &models.Admin{})
}

func AutoMigrate() {
	err := DB.AutoMigrate(&models.Asset{}, &models.User{}, &models.Accesorie{}, &models.Department{}, &models.Admin{})
>>>>>>> 2fe1552807d5a5c090e33e8e4898e3f5753702b8
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Database Migrated")
}
