package service

import (
	"GinRest/dto"
	"GinRest/entity"
	"GinRest/repository"
)

type BookService interface {
	Insert(dto dto.BookCreateDTO) entity.Book
	Update(updateDTO dto.BookUpdateDTO) entity.Book
	Delete(book entity.Book)
	All() []entity.Book
	FindByID(bookID uint64) entity.Book
	IsAllowedToEdit(userID string,bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func (service *bookService) Insert(dto dto.BookCreateDTO) entity.Book {
	panic("implement me")
}

func (service *bookService) Update(updateDTO dto.BookUpdateDTO) entity.Book {
	panic("implement me")
}

func (service *bookService) Delete(book entity.Book) {
	panic("implement me")
}

func (service *bookService) All() []entity.Book {
	panic("implement me")
}

func (service *bookService) FindByID(bookID uint64) entity.Book {
	panic("implement me")
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	panic("implement me")
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{bookRepository}
}