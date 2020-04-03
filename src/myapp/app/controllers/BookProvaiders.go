package controllers

import (
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type BookProvaider struct {
	mapper BookMapper
}

func (pr BookProvaider) GiveBooksPro() (bookspro []BookPro, err error) {
	books, err := pr.mapper.TakeBooks()
	if err != nil {
		return nil, err
	}
	for _, v := range books {
		p := BookPro{}
		p.Id = v.Id
		p.Isbn = v.Isbn
		p.BookName = v.BookName
		p.Autor = v.Autor
		p.Publisher = v.Publisher
		p.Year = v.Year
		if v.Employeeid == 1 {
			p.Status = "В наличии"
			p.Name = ""
			p.Employeeid = 0
			p.Datestart = ""
			p.Datefinish = ""
			p.Cellnumber = ""
		} else {
			p.Status = "Нет в наличии"
			p.Name = v.Name
			p.Cellnumber = strconv.Itoa(v.Cellnumber)
			p.Employeeid = v.Employeeid
			p.Datestart = v.Datestart.Format("2006-01-02")
			p.Datefinish = v.Datestart.AddDate(0, 0, 7).Format("2006-01-02")
		}

		bookspro = append(bookspro, p)
	}
	return bookspro, nil
}

func (pr BookProvaider) BookDeletePro(books IdBooks) error {
	if books.IdBook != "" {
		Id, err := strconv.Atoi(books.IdBook)
		if err != nil {
			return err
		}
		err = pr.mapper.BookDeleteOne(Id)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := pr.mapper.BookDeleteSome(books.IdBooks)
		if err != nil {
			return err
		}
		return nil
	}

}

func (pr BookProvaider) AddBookPro(bookAdd BookAdd) (err error) {
	var book BookAddPro
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
