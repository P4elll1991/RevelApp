package controllers

import (
	"time"

	_ "github.com/lib/pq"
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
