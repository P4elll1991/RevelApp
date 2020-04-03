package controllers

import (
	"time"

	_ "github.com/lib/pq"
)

type data struct {
	Books   []BookPro
	Staff   []Employee
	Journal []EventPro
}

type BookAddPro struct {
	Isbn       int
	BookName   string
	Autor      string
	Publisher  string
	Year       int
	Employeeid int
	Datestart  time.Time
}

type BookAdd struct {
	Isbn      int    `json:"isbn"`
	BookName  string `json:"bookName"`
	Autor     string `json:"autor"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
}

type IdBooks struct {
	IdBook  string
	IdBooks []int
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
