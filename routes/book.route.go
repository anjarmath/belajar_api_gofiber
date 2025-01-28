package routes

import (
	"github.com/anjarmath/01_golang_api_sederhana/controller"
	"github.com/anjarmath/01_golang_api_sederhana/middleware"
	"github.com/anjarmath/01_golang_api_sederhana/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupBookRoute(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Belajar bareng API Sederhana!")
	})

	bookRepository := repository.NewBookRepositoryPostgres()
	bookController := controller.NewBookController(bookRepository)

	book := r.Group("/book", middleware.AuthMiddleware())

	// Get Buku
	book.Get("/", bookController.GetBook)

	// Get buku berdasarkan ID
	book.Get("/:id", bookController.GetBookByID)

	// Post Buku
	book.Post("/", bookController.AddBook)

	// Update buku
	book.Patch("/:id", bookController.UpdateBook)

	// Hapus buku
	book.Delete("/:id", bookController.DeleteBook)
}
