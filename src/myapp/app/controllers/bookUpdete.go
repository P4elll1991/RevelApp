package controllers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

func (c Books) Update() revel.Result {
	var bookUpdete Book
	c.Params.BindJSON(&bookUpdete)

	err := UpdateBookPro(bookUpdete)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render()
}

func UpdateBookPro(bookUpdate Book) error {
	bookUpdate.Datestart = time.Now()
	err := UpdateBook(bookUpdate)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(b Book) error {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Добавить елемент

	connStr = "update books set isbn = $1, bookname = $2, autor = $3, publisher = $4, year = $5, Employeeid = $6,  Datestart = $7 where id = $8"
	_, err = db.Exec(connStr, b.Isbn, b.BookName, b.Autor, b.Publisher, b.Year, b.Employeeid, b.Datestart, b.Id)

	if err != nil {
		return err
	}

	return nil
}
