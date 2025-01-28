package main

import (
	"log"

	db "github.com/anjarmath/01_golang_api_sederhana/DB"
	"github.com/anjarmath/01_golang_api_sederhana/model"
	"github.com/anjarmath/01_golang_api_sederhana/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(".env is not load properly")
	}

	app := fiber.New()

	// Inisialisasi Koleksi Buku
	model.InitBook()

	// Inisialisasi Database
	db.InitDB()

	// Inisialisasi Store
	db.InitStore()

	// Tambahkan session ke fiber context
	app.Use(func(c *fiber.Ctx) error {
		sess, err := db.Store.Get(c)
		if err != nil {
			return err
		}
		c.Locals("session", sess)
		return c.Next()
	})

	// Panggil Route
	routes.SetupBookRoute(app)
	routes.SetupAuthRoute(app)

	app.Listen(":3000")
}
