package routes

import (
	"github.com/anjarmath/01_golang_api_sederhana/controller"
	"github.com/anjarmath/01_golang_api_sederhana/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoute(r *fiber.App) {

	userRepository := repository.NewUserRepository()
	authController := controller.NewAuthController(userRepository)

	auth := r.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
}
