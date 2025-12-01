package middleware

import (
	"futurisme-api/config"
	"futurisme-api/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

// AppLayerAuth adalah Layer 1 Security
// Mengecek apakah request memiliki kredensial aplikasi yang valid (X-App-Key & X-App-Secret)
func AppLayerAuth(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil header
		clientKey := c.Get("X-App-Key")
		clientSecret := c.Get("X-App-Secret")

		// Validasi dengan config dari .env
		if clientKey != cfg.Security.AppKey || clientSecret != cfg.Security.AppSecret {
			// Return 401 Unauthorized jika tidak cocok
			return response.Error(c, fiber.StatusUnauthorized, "Invalid Application Credentials", nil)
		}

		// Lanjut ke layer berikutnya
		return c.Next()
	}
}
