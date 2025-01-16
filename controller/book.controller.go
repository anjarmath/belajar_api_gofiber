package controller

import (
	"github.com/anjarmath/01_golang_api_sederhana/model"
	"github.com/anjarmath/01_golang_api_sederhana/repository"
	"github.com/gofiber/fiber/v2"
)

type bookController struct {
	bookRepository repository.BookRepository
}

func NewBookController(repo repository.BookRepository) bookController {
	return bookController{
		bookRepository: repo,
	}
}

func (bc bookController) GetBook(c *fiber.Ctx) error {
	return c.JSON(model.BooksCollection)
}

func (bc bookController) AddBook(c *fiber.Ctx) error {
	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Mohon isikan data dengan benar",
		})
	}

	valid := validateBook(book)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: "Mohon isikan data dengan benar",
		})
	}

	bc.bookRepository.AddBook(book)

	return c.JSON(fiber.Map{
		"message": "Buku berhasil ditambahkan",
		"buku":    book,
	})
}

func validateBook(book *model.Book) bool {
	return !(book.Judul == "" || book.Deskripsi == "" || book.Harga == 0)
}
