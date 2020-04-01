package controllers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/revel/revel"
)

type Staff struct {
	*revel.Controller
}

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

func (c Staff) Give() revel.Result {
	staff, err := GiveStaffPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(staff)
}

func GiveStaffPro() (staff []Employee, err error) {
	staffPro, books, err := TakeStaff()
	fmt.Println(books)
	if err != nil {
		return nil, err
	}
	for _, val := range staffPro {
		p := Employee{}
		b := []BookOfEmployee{}
		p.Id = val.Id
		p.Name = val.Name
		p.Department = val.Department
		p.Position = val.Position
		p.Cellnumber = val.Cellnumber

		for _, v := range books {
			if val.Id == v.Employeeid {
				b = append(b, v)
			}
		}

		p.Books = b
		staff = append(staff, p)
	}
	return staff, nil
}

func TakeStaff() ([]EmployeePro, []BookOfEmployee, error) {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	connStr = "SELECT * FROM staff"
	rows, err := db.Query(connStr)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	staff := []EmployeePro{}

	for rows.Next() {
		p := EmployeePro{}
		err := rows.Scan(&p.Id, &p.Name, &p.Department, &p.Position, &p.Cellnumber)
		if err != nil {
			fmt.Println(err)
			continue
		}

		staff = append(staff, p)
	}

	connStr = "SELECT id, Isbn, BookName, Employeeid, Datestart From books"
	rows, err = db.Query(connStr)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	books := []BookOfEmployee{}

	for rows.Next() {
		p := BookOfEmployee{}
		err := rows.Scan(&p.IdBook, &p.Isbn, &p.BookName, &p.Employeeid, &p.DatestartTime)
		if err != nil {
			fmt.Println(err)
			continue
		}
		p.Datestart = p.DatestartTime.Format("2006-01-02")
		p.Datefinish = p.DatestartTime.AddDate(0, 0, 7).Format("2006-01-02")

		books = append(books, p)
	}

	return staff, books, nil
}
