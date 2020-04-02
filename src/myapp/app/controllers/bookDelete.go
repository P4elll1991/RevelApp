package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/revel/revel"
)

type IdBooks struct {
	IdBook  string
	IdBooks []int
}

func (c Books) Delete() revel.Result {
	bookProvaider := BookPro{}
	var IdArr IdBooks
	IdArr.IdBook = c.Params.Query.Get("id")
	if IdArr.IdBook != "" {
		err := bookProvaider.BookDeletePro(IdArr)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		c.Params.BindJSON(&IdArr.IdBooks)

		err := bookProvaider.BookDeletePro(IdArr)
		if err != nil {
			fmt.Println(err)
		}
	}

	return c.Render()
}

func (BookPro) BookDeletePro(books IdBooks) error {
	booksMapper := Book{}
	if books.IdBook != "" {
		Id, err := strconv.Atoi(books.IdBook)
		if err != nil {
			return err
		}
		err = booksMapper.BookDeleteOne(Id)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := booksMapper.BookDeleteSome(books.IdBooks)
		if err != nil {
			return err
		}
		return nil
	}

}

func (Book) BookDeleteSome(b []int) error {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	connStr = "delete from books where id = $1"

	if err != nil {
		return err
	}
	defer db.Close()

	for _, v := range b {
		_, err = db.Exec(connStr, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Book) BookDeleteOne(id int) error {
	// Открытие базы данных

	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	connStr = "delete from books where id = $1"

	// Удаление из базы данных
	_, err = db.Exec(connStr, id)
	if err != nil {
		return err
	}

	return nil
}
