package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type BookMapper struct {
	db *sql.DB
}

func (m BookMapper) AddBook(b Book) error {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	var err error
	m.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer m.db.Close()

	// Добавить елемент

	connStr = "insert into books (isbn, bookname, autor, publisher, year, Employeeid, Datestart) values ( $1, $2, $3, $4, $5, $6, $7)"
	_, err = m.db.Exec(connStr, b.Isbn, b.BookName, b.Autor, b.Publisher, b.Year, b.Employeeid, b.Datestart)

	if err != nil {
		return err
	}

	return nil
}

func (m BookMapper) BookDelete(b []int) error {
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	connStr = "delete from books where id = $1"

	if err != nil {
		return err
	}
	defer m.db.Close()

	for _, v := range b {
		_, err = m.db.Exec(connStr, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m BookMapper) UpdateBook(b Book) error {
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer m.db.Close()

	// Добавить елемент

	connStr = "update books set isbn = $1, bookname = $2, autor = $3, publisher = $4, year = $5, Employeeid = $6,  Datestart = $7 where id = $8"
	_, err = m.db.Exec(connStr, b.Isbn, b.BookName, b.Autor, b.Publisher, b.Year, b.Employeeid, b.Datestart, b.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m BookMapper) TakeBooks() ([]Book, error) {
	// Открытие базы данных
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer m.db.Close()

	connStr = "SELECT books.id, books.isbn, books.bookname, books.autor, books.publisher, books.year, books.employeeid, books.datestart, staff.name, staff.cellnumber FROM books LEFT JOIN staff ON books.employeeid = staff.id;"
	rows, err := m.db.Query(connStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := []Book{}

	for rows.Next() {
		p := Book{}
		err := rows.Scan(&p.Id, &p.Isbn, &p.BookName, &p.Autor, &p.Publisher, &p.Year, &p.Employeeid, &p.Datestart, &p.Name, &p.Cellnumber)
		if err != nil {
			fmt.Println(err)
			continue
		}

		books = append(books, p)
	}
	return books, nil
}
