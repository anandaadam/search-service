package repositories

import (
	"search-service/models"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	SearchUsers(ctx *fiber.Ctx, input string) ([]models.User, error)
}
