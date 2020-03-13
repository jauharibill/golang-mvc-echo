package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"Person/app/Models"
)

type config struct {
	db_connection string
	db_name string
	db_host string
	db_port string
	db_username string
	db_password string
}

func Init() *gorm.DB {

	InitialDB, _ := gorm.Open(connectionMap().db_connection, assembleConfig())

	InitialDB.AutoMigrate(&Models.PersonModel{})

	return InitialDB
}

func connectionMap() config {

	conf := config{
		db_connection: os.Getenv("DB_CONNECTION"),
		db_name:     os.Getenv("DB_NAME"),
		db_host:     os.Getenv("DB_HOST"),
		db_port:     os.Getenv("DB_PORT"),
		db_username: os.Getenv("DB_USERNAME"),
		db_password: os.Getenv("DB_PASSWORD"),
	}

	return conf
}

func assembleConfig() string {

	conf := connectionMap().db_username + ":" +
		connectionMap().db_password + "@(" +
		connectionMap().db_host + ":" +
		connectionMap().db_port + ")/" +
		connectionMap().db_name + "?" +
		"parseTime=true"

	return conf
}
