package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type BookMapper struct {
	db *sql.DB
}

// метод добавдяющий книгу в БД

func (m BookMapper) AddBook(b Book) error {
	var err error
	m.db, err = InitDB() // Инициализация БД
	if err != nil {
		return err
	}

	// Добавить елемент
	connStr := "insert into books (isbn, bookname, autor, publisher, year, Employeeid, Datestart) values ( $1, $2, $3, $4, $5, $6, $7)"
	_, err = m.db.Exec(connStr, b.Isbn, b.BookName, b.Autor, b.Publisher, b.Year, b.Employeeid, b.Datestart)

	if err != nil {
		return err
	}

	return nil
}

// Метод удаляющий книги из Бд

func (m BookMapper) BookDelete(b []int) error {
	var err error
	m.db, err = InitDB() // Инициализация
	if err != nil {
		return err
	}
	connStr := "delete from books where id = $1"

	// Перебор среза id
	for _, v := range b {
		_, err = m.db.Exec(connStr, v) // удаление строки
		if err != nil {
			return err
		}
	}
	return nil
}

// метод обновления данных в БД

func (m BookMapper) UpdateBook(b Book) error {
	var err error
	m.db, err = InitDB() // Инициализация
	if err != nil {
		return err
	}

	// Обновить елемент

	connStr := "update books set isbn = $1, bookname = $2, autor = $3, publisher = $4, year = $5, Employeeid = $6,  Datestart = $7 where id = $8"
	_, err = m.db.Exec(connStr, b.Isbn, b.BookName, b.Autor, b.Publisher, b.Year, b.Employeeid, b.Datestart, b.Id)

	if err != nil {
		return err
	}

	return nil
}

// Получение данных из БД

func (m BookMapper) TakeBooks() ([]Book, error) {

	var err error
	m.db, err = InitDB() // Инициализация
	if err != nil {
		return nil, err
	}

	connStr := "SELECT books.id, books.isbn, books.bookname, books.autor, books.publisher, books.year, books.employeeid, books.datestart, staff.name, staff.cellnumber FROM books LEFT JOIN staff ON books.employeeid = staff.id;"
	rows, err := m.db.Query(connStr) // Запрос
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := []Book{}

	for rows.Next() { // Перевод данных в структуру
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
