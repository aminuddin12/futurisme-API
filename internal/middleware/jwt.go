package middleware

import (
	"strings"

	"futurisme-api/config"
	jwtUtil "futurisme-api/pkg/utils/jwt"
	"futurisme-api/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTProtected adalah Middleware Layer 2
// Memvalidasi Bearer Token dan menyimpan Claims ke Context
func JWTProtected(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Ambil Header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, fiber.StatusUnauthorized, "Missing Authorization Header", nil)
		}

		// 2. Cek Format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid Token Format", nil)
		}

		tokenString := parts[1]

		// 3. Parse Token
		claims := &jwtUtil.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.Security.JWTSecret), nil
		})

		// 4. Validasi Error Parsing
		if err != nil {
			if err == jwt.ErrTokenExpired {
				return response.Error(c, fiber.StatusUnauthorized, "Token Expired", nil)
			}
			return response.Error(c, fiber.StatusUnauthorized, "Invalid Token", nil)
		}

		// 5. Validasi Token Valid
		if !token.Valid {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid Token", nil)
		}

		// 6. Simpan Data User ke Context (Local Storage Request)
		// Agar bisa diakses di Handler selanjutnya (misal: c.Locals("user_id"))
		c.Locals("user_id", claims.UserID)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}
