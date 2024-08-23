package handlers

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}

// Response function for creating a new response and sending it to the client.
func Response(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := BaseResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	return c.Status(statusCode).JSON(response)
}
