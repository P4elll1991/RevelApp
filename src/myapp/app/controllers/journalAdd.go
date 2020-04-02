package controllers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/revel/revel"
)

func (c Journal) Add() revel.Result {
	journalProvaider := EventPro{}
	var event Event
	c.Params.BindJSON(&event)

	err := journalProvaider.AddEventPro(event)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render()
}

func (EventPro) AddEventPro(event Event) error {
	journalMapper := Event{}
	event.DateEvent = time.Now()

	err := journalMapper.AddEvent(event)
	if err != nil {
		return err
	}

	return nil
}

func (Event) AddEvent(event Event) error {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Добавить елемент

	connStr = "insert into journal (event, bookid, isbn, BookName, employeeid, Name, Cellnumber, dateevent) values ( $1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = db.Exec(connStr, event.Event, event.BookId, event.Isbn, event.BookName, event.EmployeeId, event.Name, event.Cellnumber, event.DateEvent)

	if err != nil {
		return err
	}

	return nil
}
