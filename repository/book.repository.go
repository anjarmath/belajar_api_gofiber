package repository

import "github.com/anjarmath/01_golang_api_sederhana/model"

type BookRepository interface {
	AddBook(book *model.Book)
}
