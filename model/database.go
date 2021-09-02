package model

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn     *gorm.DB
	dsnPattern = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s"
)

func InitDatabase(conf DBConf) {
	var err error
	dsn := fmt.Sprintf(dsnPattern, conf.Host, conf.User, conf.Pass, conf.Name, conf.Port, conf.TZ)
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
		&operation{},
		&assets{},
		&category{},
		&budgetPosition{},
	)
	if err != nil {
		panic(fmt.Sprintf("cannot migrate: %s", err.Error()))
	}
	log.Println("Auto migration succeeded")
}

type DBConf struct {
	Host string
	Port string
	User string
	Pass string
	Name string
	TZ   string
}
