package repository

import (
	db "github.com/anjarmath/01_golang_api_sederhana/DB"
	"github.com/anjarmath/01_golang_api_sederhana/model"
)

type bookRepositoryImplPostgres struct{}

// AddBook implements BookRepository.
func (b bookRepositoryImplPostgres) AddBook(book *model.Book) error {
	err := db.DB.Create(book).Error
	return err
}

// DeleteBook implements BookRepository.
func (b bookRepositoryImplPostgres) DeleteBook(id string) error {
	err := db.DB.Delete(&model.Book{}, "id = ?", id).Error
	return err
}

// GetBookByID implements BookRepository.
func (b bookRepositoryImplPostgres) GetBookByID(id string) (*model.Book, error) {
	book := new(model.Book)
	if err := db.DB.Preload("InterestedUsers").First(book, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// GetBooks implements BookRepository.
func (b bookRepositoryImplPostgres) GetBooks() (*[]model.Book, error) {
	books := new([]model.Book)
	if err := db.DB.Preload("InterestedUsers").Find(books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// UpdateBook implements BookRepository.
func (b bookRepositoryImplPostgres) UpdateBook(id string, book *model.Book) error {
	err := db.DB.Model(&model.Book{}).Where("id = ?", id).Updates(book).Error
	return err
}

func NewBookRepositoryPostgres() BookRepository {
	return bookRepositoryImplPostgres{}
}
