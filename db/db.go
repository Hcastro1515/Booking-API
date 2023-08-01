package db

import (
	"log"

	"github.com/hcastro1515/booking-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Booking{}, &models.Service{})

	return db
}
