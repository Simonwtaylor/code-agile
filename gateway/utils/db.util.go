package utils

import (
	"fmt"
	"log"

	e "github.com/Simonwtaylor/code-agile/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Unable to connect to DB", err.Error())
		return nil, err
	}

	log.Println("Successfully connected to DB 🔌")

	err = db.AutoMigrate(&e.TaskEntity{}, &e.TaskStatusEntity{})

	if err != nil {
		fmt.Println("Unable to migrate entities", err.Error())
		return nil, err
	}

	log.Println("Entities Migrated 🏁")

	return db, err
}
