package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	DBConn     *gorm.DB
	dsnPattern = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw"
)

func InitDatabase() {
	var err error
	dsn := fmt.Sprintf(dsnPattern, "localhost", "postgres", os.Getenv("DB_PASS"), "postgres", "5432")
	for {
		DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	log.Println("Connection Opened to Database")

	err = DBConn.AutoMigrate(
		&Operation{},
		&Assets{},
		&Category{},
	)
	if err != nil {
		panic(fmt.Sprintf("cannot migrate: %s", err.Error()))
	}
	log.Println("Auto migration succeeded")
}
