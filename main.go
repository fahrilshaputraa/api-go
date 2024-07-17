package main

import (
	"log"

	"apigo/database"
	_ "apigo/docs"

	"apigo/lib"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 8Lqg5@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api/v1
func main() {
	app := fiber.New()
	db := database.ConnectDB(app)
	if db == nil {
		log.Fatal("Failed to connect to database")
	}
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	lib.Router(app, db)

	log.Fatal(app.Listen(":3000"))
}
