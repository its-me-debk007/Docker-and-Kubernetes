package database

import (
	"github.com/its-me-debk007/Docker-and-Kubernetes/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	if err := db.AutoMigrate(
		new(model.User),
	); err != nil {
		log.Fatalln("AUTO_MIGRATION_ERROR")
	}

}
