package repository

import "github.com/anjarmath/01_golang_api_sederhana/model"

type bookRepositoryImpl struct{}

// DeleteBook implements BookRepository.
func (b bookRepositoryImpl) DeleteBook(id string) error {
	panic("unimplemented")
}

// GetBookByID implements BookRepository.
func (b bookRepositoryImpl) GetBookByID(id string) (*model.Book, error) {
	panic("unimplemented")
}

// GetBooks implements BookRepository.
func (b bookRepositoryImpl) GetBooks() (*[]model.Book, error) {
	return &model.BooksCollection, nil
}

// UpdateBook implements BookRepository.
func (b bookRepositoryImpl) UpdateBook(id string, book *model.Book) error {
	panic("unimplemented")
}

// AddBook implements BookRepository.
func (b bookRepositoryImpl) AddBook(book *model.Book) error {
	model.BooksCollection = append(model.BooksCollection, *book)
	return nil
}

// Membuat repository
func NewBookRepository() BookRepository {
	return bookRepositoryImpl{}
}
