package controllers

import (
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	Id         int
	Name       string
	Department string
	Position   string
	Cellnumber int
	Books      []BookOfEmployee
}

type BookOfEmployee struct {
	IdBook        int
	Isbn          int
	BookName      string
	Employeeid    int
	DatestartTime time.Time
	Datestart     string
	Datefinish    string
}

type EmployeePro struct {
	Id         int
	Name       string
	Department string
	Position   string
	Cellnumber int
}

type IdStaff struct {
	IdEmp   string
	IdStaff []int
}
