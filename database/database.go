package database

import (
	"apigo/lib"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(app *fiber.App) *gorm.DB {
	dsn := "host=172.17.0.4 user=fahrilshaputra password=fahrilshaputra000 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&lib.User{})
	if err != nil {
		panic(err)
	}
	return db
}
