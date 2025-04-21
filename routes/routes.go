package routes

import (
	"github.com/gofiber/fiber/v2"
	c "go-workshop/controllers"
	m "go-workshop/middlewares"
)

func Routes(app *fiber.App) {

	// api/v1
	api := app.Group("/api")
	v1 := api.Group("/v1")

	public := v1.Group("/users")
	public.Get("", c.GetUsers)
	public.Get("/filter", c.GetUser)
	public.Get("/gen", c.GetUsersGen)
	public.Get("/find", c.FindUser)

	protect := v1.Group("/users", m.BasicAuth())
	protect.Post("/", c.AddUser)
	protect.Put("/:id", c.UpdateUser)
	protect.Delete("/:id", c.RemoveUser)
}
