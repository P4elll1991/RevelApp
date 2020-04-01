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
	var IdArr IdBooks
	IdArr.IdBook = c.Params.Query.Get("id")
	if IdArr.IdBook != "" {
		err := BookDeletePro(IdArr)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		c.Params.BindJSON(&IdArr.IdBooks)

		err := BookDeletePro(IdArr)
		if err != nil {
			fmt.Println(err)
		}
	}

	return c.Render()
}

func BookDeletePro(books IdBooks) error {
	if books.IdBook != "" {
		Id, err := strconv.Atoi(books.IdBook)
		if err != nil {
			return err
		}
		err = BookDelete1(Id)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := BookDelete2(books.IdBooks)
		if err != nil {
			return err
		}
		return nil
	}

}

func BookDelete2(b []int) error {
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

func BookDelete1(id int) error {
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
