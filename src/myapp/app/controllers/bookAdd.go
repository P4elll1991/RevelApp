package controllers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/revel/revel"
)

type BookAdd struct {
	Isbn      int    `json:"isbn"`
	BookName  string `json:"bookName"`
	Autor     string `json:"autor"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
}

type BookAddPro struct {
	Isbn       int
	BookName   string
	Autor      string
	Publisher  string
	Year       int
	Employeeid int
	Datestart  time.Time
}

func (c Books) Add() revel.Result {
	var bookAdd BookAdd
	c.Params.BindJSON(&bookAdd)

	err := AddBookPro(bookAdd)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render()
}

func AddBookPro(bookAdd BookAdd) (err error) {
	var book BookAddPro
	book.Isbn = bookAdd.Isbn
	book.BookName = bookAdd.BookName
	book.Autor = bookAdd.Autor
	book.Publisher = bookAdd.Publisher
	book.Year = bookAdd.Year
	book.Employeeid = 1
	book.Datestart = time.Now()
	err = AddBook(book)
	if err != nil {
		return err
	}
	return nil
}

func AddBook(b BookAddPro) error {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Добавить елемент

	connStr = "insert into books (isbn, bookname, autor, publisher, year, Employeeid, Datestart) values ( $1, $2, $3, $4, $5, $6, $7)"
	_, err = db.Exec(connStr, b.Isbn, b.BookName, b.Autor, b.Publisher, b.Year, b.Employeeid, b.Datestart)

	if err != nil {
		return err
	}

	return nil
}
