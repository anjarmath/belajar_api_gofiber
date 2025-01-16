package main

import (
	"github.com/anjarmath/01_golang_api_sederhana/model"
	"github.com/anjarmath/01_golang_api_sederhana/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inisialisasi Koleksi Buku
	model.InitBook()

	// Panggil Route
	routes.SetupBookRoute(app)

	app.Listen(":3000")
}
