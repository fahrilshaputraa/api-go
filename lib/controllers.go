package lib

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetUsers mengambil semua pengguna dari database dan mengembalikannya dalam format JSON.
// @Summary Dapatkan semua pengguna
// @Description Mengambil daftar semua pengguna yang ada di database
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func GetUsers(c *fiber.Ctx, db *gorm.DB) error {
	var users []User
	result := db.Find(&users)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pengguna"})
	}

	return c.JSON(users)
}
