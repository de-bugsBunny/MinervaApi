package presenter

import (
	"minerva_api/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

// User is the presenter object which will be taken in the request by Handler
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhotoUrl string `json:"photo_url"`
}

// UserSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func UserSuccessResponse(data *entities.User) *fiber.Map {

	newUser := User{
		Name:     data.Name,
		Email:    data.Email,
		PhotoUrl: data.PhotoUrl,
	}
	return &fiber.Map{
		"status": true,
		"data":   newUser,
		"error":  nil,
	}
}

// TopicErrorResponse is the ErrorResponse that will be passed in the response by Handler
func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
