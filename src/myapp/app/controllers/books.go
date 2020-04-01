package controllers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

type Books struct {
	*revel.Controller
}

type BookPro struct {
	Id         int
	Isbn       int
	BookName   string
	Autor      string
	Publisher  string
	Year       int
	Status     string
	Name       string
	Employeeid int
	Datestart  string
	Datefinish string
}

type Book struct {
	Id         int
	Isbn       int
	BookName   string
	Autor      string
	Publisher  string
	Year       int
	Employeeid int
	Name       string
	Datestart  time.Time
}

func (c Books) Give() revel.Result {
	allbooks, err := GiveBooksPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(allbooks)
}

func GiveBooksPro() (bookspro []BookPro, err error) {
	books, err := TakeBooks()
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
		} else {
			p.Status = "Нет в наличии"
			p.Name = v.Name
			p.Employeeid = v.Employeeid
			p.Datestart = v.Datestart.Format("2006-01-02")
			p.Datefinish = v.Datestart.AddDate(0, 0, 7).Format("2006-01-02")
		}

		bookspro = append(bookspro, p)
	}
	return bookspro, nil
}
func TakeBooks() ([]Book, error) {
	// Открытие базы данных

	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	connStr = "SELECT books.id, books.isbn, books.bookname, books.autor, books.publisher, books.year, books.employeeid, books.datestart, staff.name FROM books LEFT JOIN staff ON books.employeeid = staff.id;"
	rows, err := db.Query(connStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := []Book{}

	for rows.Next() {
		p := Book{}
		err := rows.Scan(&p.Id, &p.Isbn, &p.BookName, &p.Autor, &p.Publisher, &p.Year, &p.Employeeid, &p.Datestart, &p.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		books = append(books, p)
	}
	return books, nil
}
