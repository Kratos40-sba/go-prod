package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	logLib "github.com/sirupsen/logrus"
	"os"
)

func NewDatabase() (*gorm.DB, error) {
	logLib.Info("Setting up new database connection")
	dbUserName := os.Getenv("DB_UNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUserName, dbTable, dbPassword)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}
	if err := db.DB().Ping(); err != nil {
		return db, err
	}
	return db, nil
}
