package http

import (
	"futurisme-api/internal/modules/user/usecase"
	"futurisme-api/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// GetProfile menangani GET /users/profile
// Endpoint ini dilindungi JWT Middleware
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	// 1. Ambil UserID dari Context (hasil set dari Middleware JWT)
	userID := c.Locals("user_id").(uint)

	// 2. Panggil Business Logic
	user, err := h.userUseCase.GetProfile(c.Context(), userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	if user == nil {
		return response.Error(c, fiber.StatusNotFound, "User not found", nil)
	}

	return response.Success(c, fiber.StatusOK, "User profile fetched successfully", user)
}
