package db

import (
	"encoding/json"
	//"fmt"
	"log"
	"os"
	//"TestTask/models"

	//"gorm.os/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    User     string `json:"user"`
    Password string `json:"password"`
    DB       string `json:"db"`
}

var DB *gorm.DB

func InitDB() {
    config := Config{}
    configFile, err := os.Open("config.json")
    if err != nil {
        log.Fatal(err)
    }
    defer configFile.Close()

    err = json.NewDecoder(configFile).Decode(&config)
    if err != nil {
        log.Fatal(err)
    }

    /*dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.Host, config.Port, config.User, config.Password, config.DB)
    var errDB error
    DB, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if errDB != nil {
        log.Fatal(errDB)
    }

    err = DB.AutoMigrate(&models.Building{})
    if err != nil {
        log.Fatal(err)
    }*/
}