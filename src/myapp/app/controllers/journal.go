package controllers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

type EventPro struct {
	Id         int
	Event      string
	BookId     int
	DateEvent  time.Time
	EmployeeId int
}

type Event struct {
	Id          int
	Event       string
	BookId      int
	BookNameJ   string
	IsbnJ       int
	DateEvent   string
	EmployeeId  int
	NameJ       string
	CellnumberJ int
}

type Journal struct {
	*revel.Controller
}

func (c Journal) Give() revel.Result {
	Journal, err := GiveJournalPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(Journal)
}

func GiveJournalPro() (journal []Event, err error) {

	journalPro, books, staff, err := TakeJournal()
	if err != nil {
		return nil, err
	}

	for _, val := range journalPro {
		p := Event{}
		for _, v := range books {
			if val.BookId == v.Id {
				p.BookNameJ = v.BookName
				p.IsbnJ = v.Isbn
			}
		}
		for _, v := range staff {
			if val.EmployeeId == v.Id {
				p.NameJ = v.Name
				p.CellnumberJ = v.Cellnumber
			}
		}

		p.Id = val.Id
		p.Event = val.Event
		p.BookId = val.BookId
		p.EmployeeId = val.EmployeeId
		p.DateEvent = val.DateEvent.Format("2006-01-02")

		journal = append(journal, p)
	}

	return journal, nil
}

func TakeJournal() (journal []EventPro, books []Book, staff []Employee, err error) {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, nil, err
	}
	defer db.Close()

	connStr = "SELECT id, Name, Cellnumber FROM staff"
	rows, err := db.Query(connStr)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()
	staff = []Employee{}

	for rows.Next() {
		p := Employee{}
		err := rows.Scan(&p.Id, &p.Name, &p.Cellnumber)
		if err != nil {
			fmt.Println(err)
			continue
		}

		staff = append(staff, p)
	}

	connStr = "SELECT id, Isbn, BookName From books"
	rows, err = db.Query(connStr)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()
	books = []Book{}

	for rows.Next() {
		p := Book{}
		err := rows.Scan(&p.Id, &p.Isbn, &p.BookName)
		if err != nil {
			return nil, nil, nil, err
		}

		books = append(books, p)
	}

	connStr = "SELECT * From journal"
	rows, err = db.Query(connStr)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()
	journal = []EventPro{}

	for rows.Next() {
		p := EventPro{}
		err := rows.Scan(&p.Id, &p.Event, &p.BookId, &p.EmployeeId, &p.DateEvent)
		if err != nil {
			return nil, nil, nil, err
		}

		journal = append(journal, p)
	}

	return journal, books, staff, nil
}
