package migration

import (
	"fmt"
	"go-miniproject-iqbalikhlasulamal/database"
	"go-miniproject-iqbalikhlasulamal/models/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Scholarship{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
