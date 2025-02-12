package controller

import (
	"github.com/anjarmath/01_golang_api_sederhana/model"
	"github.com/anjarmath/01_golang_api_sederhana/repository"
	"github.com/anjarmath/01_golang_api_sederhana/utils"
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
	books, err := bc.bookRepository.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Error{
			Message: "Data tidak ditemukan",
		})
	}
	return c.JSON(books)
}

func (bc bookController) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := bc.bookRepository.GetBookByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Error{
			Message: "Data tidak ditemukan",
		})
	}
	return c.JSON(book)
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

	id := utils.GetIdFromSession(c)
	book.UserId = id

	err := bc.bookRepository.AddBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: "Terjadi kesalahan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Buku berhasil ditambahkan",
		"buku":    book,
	})
}

func (bc bookController) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	reqUpdate := new(model.Book)
	if err := c.BodyParser(reqUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Mohon isikan data dengan benar",
		})
	}

	err := bc.bookRepository.UpdateBook(id, reqUpdate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: "Terjadi kesalahan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Buku berhasil diubah",
		"buku":    reqUpdate,
	})
}

func (bc bookController) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	err := bc.bookRepository.DeleteBook(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: "Terjadi kesalahan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Buku berhasil dihapus",
	})
}

func validateBook(book *model.Book) bool {
	return !(book.Judul == "" || book.Deskripsi == "" || book.Harga == 0)
}
