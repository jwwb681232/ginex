package service

import (
	"GinRest/dto"
	"GinRest/entity"
	"GinRest/repository"
	"fmt"
	"github.com/mashingan/smapping"
	"log"
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
	book := entity.Book{}
	err := smapping.FillStruct(&book,smapping.MapFields(&dto))
	if err != nil {
		log.Fatalf("Failed to map %v",err)
	}

	res := service.bookRepository.CreateBook(book)
	return res
}

func (service *bookService) Update(dto dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book,smapping.MapFields(&dto))
	if err != nil {
		log.Fatalf("Failed to map %v",err)
	}

	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service *bookService) Delete(book entity.Book) {
	service.bookRepository.DeleteBook(book)
}

func (service *bookService) All() []entity.Book {
	return service.bookRepository.Books()
}

func (service *bookService) FindByID(bookID uint64) entity.Book {
	return service.bookRepository.FindBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	book := service.bookRepository.FindBookByID(bookID)
	id := fmt.Sprintf("%v",book.UserID)

	return userID == id
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{bookRepository}
}