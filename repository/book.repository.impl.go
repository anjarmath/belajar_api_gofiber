package repository

import "github.com/anjarmath/01_golang_api_sederhana/model"

type bookRepositoryImpl struct{}

// AddBook implements BookRepository.
func (b bookRepositoryImpl) AddBook(book *model.Book) {
	model.BooksCollection = append(model.BooksCollection, *book)
}

// Membuat repository
func NewBookRepository() BookRepository {
	return bookRepositoryImpl{}
}
