package repository

import "github.com/anjarmath/01_golang_api_sederhana/model"

type BookRepository interface {
	GetBooks() (*[]model.Book, error)
	GetBookByID(id string) (*model.Book, error)
	AddBook(book *model.Book) error
	UpdateBook(id string, book *model.Book) error
	DeleteBook(id string) error
}
