package controllers

import (
	"time"

	_ "github.com/lib/pq"
)

type Event struct {
	Id          int
	Event       string
	BookId      int
	IsbnJ       int
	BookNameJ   string
	DateEvent   time.Time
	EmployeeId  int
	NameJ       string
	CellnumberJ int
}
