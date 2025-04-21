package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuth() func (*fiber.Ctx) error {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	})
}