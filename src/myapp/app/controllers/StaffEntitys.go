package controllers

import (
	_ "github.com/lib/pq"
)

type Employee struct {
	Id         int
	Name       string
	Department string
	Position   string
	Cellnumber int
	Books      []Book
}
