package controllers

import (
	"time"

	_ "github.com/lib/pq"
)

type data struct {
	Books   []Book
	Staff   []Employee
	Journal []Event
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
