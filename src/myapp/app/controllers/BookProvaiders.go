package controllers

import (
	"time"

	_ "github.com/lib/pq"
)

type BookProvaider struct {
	mapper BookMapper
}

//метод возвращающий данные книг

func (pr BookProvaider) GiveBooksPro() (books []Book, err error) {
	books, err = pr.mapper.TakeBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

//метод удаляющий книги

func (pr BookProvaider) BookDeletePro(books []int) error {

	err := pr.mapper.BookDelete(books)
	if err != nil {
		return err
	}
	return nil

}

//метод добавляющий книгу

func (pr BookProvaider) AddBookPro(book Book) (err error) {
	book.Employeeid = 1
	book.Datestart = time.Now()
	err = pr.mapper.AddBook(book)
	if err != nil {
		return err
	}
	return nil
}

// метод обновляющий данные о книгах

func (pr BookProvaider) UpdateBookPro(bookUpdate Book) error {

	bookUpdate.Datestart = time.Now()
	err := pr.mapper.UpdateBook(bookUpdate)
	if err != nil {
		return err
	}
	return nil
}
