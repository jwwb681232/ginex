package repository

import (
	"GinRest/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book entity.Book) entity.Book
	UpdateBook(book entity.Book) entity.Book
	DeleteBook(book entity.Book)
	Books() []entity.Book
	FindBookByID(bookID uint64) entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

func (db *bookConnection) CreateBook(book entity.Book) entity.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}

func (db *bookConnection) UpdateBook(book entity.Book) entity.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}

func (db *bookConnection) DeleteBook(book entity.Book) {
	db.connection.Delete(&book)
}

func (db *bookConnection) Books() []entity.Book {
	var books []entity.Book
	db.connection.Preload("User").Find(&books)
	return books
}

func (db *bookConnection) FindBookByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book,bookID)
	return book
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{connection: db}
}

