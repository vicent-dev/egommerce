package app

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (s *server) database() {
	var err error

	dbUser := os.Getenv("DB_USERNAME")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connection := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	s.db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
