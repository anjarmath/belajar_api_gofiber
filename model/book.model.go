package model

type Book struct {
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	Harga     int    `json:"harga"`
}

var BooksCollection []Book

func InitBook() {
	BooksCollection = []Book{}
}
