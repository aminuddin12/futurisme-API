package http

import (
	"futurisme-api/internal/modules/auth/usecase"
	"futurisme-api/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUseCase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// Register menangani POST /auth/register
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req usecase.RegisterRequest

	// Parsing Body
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	// Validasi Sederhana
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return response.Error(c, fiber.StatusBadRequest, "Name, Email, and Password are required", nil)
	}

	// Panggil UseCase
	user, err := h.authUseCase.Register(c.Context(), &req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.Success(c, fiber.StatusCreated, "User registered successfully", user)
}

// Login menangani POST /auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req usecase.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if req.Email == "" || req.Password == "" {
		return response.Error(c, fiber.StatusBadRequest, "Email and Password are required", nil)
	}

	loginResp, err := h.authUseCase.Login(c.Context(), &req)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.Success(c, fiber.StatusOK, "Login successful", loginResp)
}
