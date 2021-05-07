package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wambugucoder/simple-go-service/models"
	"log"
	"strconv"
)

func Connect() {
	var err error

	portString := ExtractEnvKey("PSQL_PORT")
	port, err := strconv.ParseUint(portString, 10, 32)

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", ExtractEnvKey("PSQL_HOST"), port,
		ExtractEnvKey("PSQL_USER"), ExtractEnvKey("PSQL_PASS"), ExtractEnvKey("PSQL_DB")))

	if err != nil {
		panic("Failed To Connect To Db")
	}
	DB.LogMode(true)
	log.Println("Connected To Database")

	//MIGRATE
	DB.DropTableIfExists(&models.User{}).AutoMigrate(&models.User{})

	//
	log.Println("Database has been migrated")
}
