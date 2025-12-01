package response

import "github.com/gofiber/fiber/v2"

// WebResponse adalah format standar output JSON
type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Success mengembalikan respon sukses 200/201
func Success(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(WebResponse{
		Code:    code,
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Error mengembalikan respon error (400, 404, 500, dll)
func Error(c *fiber.Ctx, code int, message string, err interface{}) error {
	return c.Status(code).JSON(WebResponse{
		Code:    code,
		Status:  "error",
		Message: message,
		Error:   err,
	})
}
