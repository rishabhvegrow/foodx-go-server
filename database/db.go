package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "fmt"
    "os"

)

var db *gorm.DB

func ConnectDB() {
    dbUser := os.Getenv("DB_USER")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbDatabse := os.Getenv("DB_DATABASE")
    dbPassword := os.Getenv("DB_PASSWORD")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbDatabse)
    // dsn := "root:password@tcp(127.0.0.1:3306)/foodX?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("Failed to connect to the database: " + err.Error())
    }
}

func GetDB() (*gorm.DB) {
    return db
}
