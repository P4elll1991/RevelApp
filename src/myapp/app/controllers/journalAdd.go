package controllers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/revel/revel"
)

func (c Journal) Add() revel.Result {
	var event EventPro
	c.Params.BindJSON(&event)

	err := AddEventPro(event)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render()
}

func AddEventPro(event EventPro) error {

	event.DateEvent = time.Now()

	err := AddEvent(event)
	if err != nil {
		return err
	}

	return nil
}

func AddEvent(event EventPro) error {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Добавить елемент

	connStr = "insert into journal (event, bookid, employeeid, dateevent) values ( $1, $2, $3, $4)"
	_, err = db.Exec(connStr, event.Event, event.BookId, event.EmployeeId, event.DateEvent)

	if err != nil {
		return err
	}

	return nil
}
