package controllers

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

type Event struct {
	Id         int
	Event      string
	BookId     int
	Isbn       int
	BookName   string
	DateEvent  time.Time
	EmployeeId int
	Name       string
	Cellnumber int
}

type EventPro struct {
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
	journalProvaider := EventPro{}
	Journal, err := journalProvaider.GiveJournalPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(Journal)
}

func (EventPro) GiveJournalPro() (journal []EventPro, err error) {
	journalMapper := Event{}

	journalPro, err := journalMapper.TakeJournal()
	if err != nil {
		return nil, err
	}

	for _, val := range journalPro {
		p := EventPro{}

		p.BookNameJ = val.BookName
		p.IsbnJ = val.Isbn
		p.NameJ = val.Name
		p.CellnumberJ = val.Cellnumber
		p.Id = val.Id
		p.Event = val.Event
		p.BookId = val.BookId
		p.EmployeeId = val.EmployeeId
		p.DateEvent = val.DateEvent.Format("2006-01-02")

		journal = append(journal, p)
	}

	return journal, nil
}

func (Event) TakeJournal() (journal []Event, err error) {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	connStr = "SELECT * From journal"
	rows, err := db.Query(connStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	journal = []Event{}

	for rows.Next() {
		p := Event{}
		err := rows.Scan(&p.Id, &p.Event, &p.BookId, &p.EmployeeId, &p.DateEvent, &p.Isbn, &p.BookName, &p.Name, &p.Cellnumber)
		if err != nil {
			return nil, err
		}

		journal = append(journal, p)
	}

	return journal, nil
}
