package lib

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) {
	v1 := app.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.Get("/", func(c *fiber.Ctx) error {
				return GetUsers(c, db)
			})
		}
	}
}
