package controllers

import (
	"time"

	_ "github.com/lib/pq"
)

type BookProvaider struct {
	mapper BookMapper
}

func (pr BookProvaider) GiveBooksPro() (books []Book, err error) {
	books, err = pr.mapper.TakeBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (pr BookProvaider) BookDeletePro(books []int) error {

	err := pr.mapper.BookDelete(books)
	if err != nil {
		return err
	}
	return nil

}

func (pr BookProvaider) AddBookPro(bookAdd Book) (err error) {
	var book Book
	book.Isbn = bookAdd.Isbn
	book.BookName = bookAdd.BookName
	book.Autor = bookAdd.Autor
	book.Publisher = bookAdd.Publisher
	book.Year = bookAdd.Year
	book.Employeeid = 1
	book.Datestart = time.Now()
	err = pr.mapper.AddBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (pr BookProvaider) UpdateBookPro(bookUpdate Book) error {

	bookUpdate.Datestart = time.Now()
	err := pr.mapper.UpdateBook(bookUpdate)
	if err != nil {
		return err
	}
	return nil
}
