package routes

import (
	"github.com/anjarmath/01_golang_api_sederhana/controller"
	"github.com/anjarmath/01_golang_api_sederhana/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupBookRoute(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Belajar bareng API Sederhana!")
	})

	bookRepository := repository.NewBookRepositoryPostgres()
	bookController := controller.NewBookController(bookRepository)

	// Get Buku
	r.Get("/book", bookController.GetBook)

	// Get buku berdasarkan ID
	r.Get("/book/:id", bookController.GetBookByID)

	// Post Buku
	r.Post("/book", bookController.AddBook)

	// Update buku
	r.Patch("/book/:id", bookController.UpdateBook)

	// Hapus buku
	r.Delete("/book/:id", bookController.DeleteBook)
}
