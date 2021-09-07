package routes

import (
	"api/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(r *fiber.App) {
	r.Get("/", controller.Home)
	r.Post("/register", controller.Register)
	r.Post("/login", controller.Login)
	r.Get("/user", controller.User)
	r.Post("/logout", controller.Logout)
}
