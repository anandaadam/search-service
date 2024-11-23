package repositories

import (
	"search-service/models"

	"github.com/gofiber/fiber/v2"
)

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (userRepository *UserRepositoryImpl) SearchUsers(ctx *fiber.Ctx, input string) ([]models.User, error) {

}
