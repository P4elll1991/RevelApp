package controllers

import (
	"database/sql"
	"fmt"
	"strconv"
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
	Cellnumber string
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
	Cellnumber int
	Datestart  time.Time
}

func (c Books) Give() revel.Result {
	bookProvaider := BookPro{}
	allbooks, err := bookProvaider.GiveBooksPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(allbooks)
}

func (BookPro) GiveBooksPro() (bookspro []BookPro, err error) {
	booksMapper := Book{}
	books, err := booksMapper.TakeBooks()
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
func (Book) TakeBooks() ([]Book, error) {
	// Открытие базы данных

	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	connStr = "SELECT books.id, books.isbn, books.bookname, books.autor, books.publisher, books.year, books.employeeid, books.datestart, staff.name, staff.cellnumber FROM books LEFT JOIN staff ON books.employeeid = staff.id;"
	rows, err := db.Query(connStr)
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
