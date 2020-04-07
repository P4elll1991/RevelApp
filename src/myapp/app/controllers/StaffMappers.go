package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type StaffMapper struct {
	db *sql.DB
}

func (m StaffMapper) UpStaff(staff Employee) error {
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer m.db.Close()

	// Добавить елемент

	connStr = "update staff set name = $1, department = $2, position = $3, cellnumber = $4 where id = $5"
	_, err = m.db.Exec(connStr, staff.Name, staff.Department, staff.Position, staff.Cellnumber, staff.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m StaffMapper) TakeStaff() ([]Employee, []Book, error) {
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, err
	}
	defer m.db.Close()

	connStr = "SELECT * FROM staff 	WHERE id != 1"
	rows, err := m.db.Query(connStr)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	staff := []Employee{}

	for rows.Next() {
		p := Employee{}
		err := rows.Scan(&p.Id, &p.Name, &p.Department, &p.Position, &p.Cellnumber)
		if err != nil {
			fmt.Println(err)
			continue
		}

		staff = append(staff, p)
	}

	connStr = "SELECT id, Isbn, BookName, Employeeid, Datestart From books"
	rows, err = m.db.Query(connStr)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	books := []Book{}

	for rows.Next() {
		p := Book{}
		err := rows.Scan(&p.Id, &p.Isbn, &p.BookName, &p.Employeeid, &p.Datestart)
		if err != nil {
			fmt.Println(err)
			continue
		}

		books = append(books, p)
	}

	return staff, books, nil
}

func (m StaffMapper) AddStaff(staff Employee) error {
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer m.db.Close()

	// Добавить елемент

	connStr = "insert into staff (name, department, position, cellnumber) values ( $1, $2, $3, $4)"
	_, err = m.db.Exec(connStr, staff.Name, staff.Department, staff.Position, staff.Cellnumber)

	if err != nil {
		return err
	}

	return nil
}

func (m StaffMapper) StaffDelete(s []int) error {
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	connStr = "delete from staff where id = $1"

	if err != nil {
		return err
	}
	defer m.db.Close()

	for _, v := range s {
		_, err = m.db.Exec(connStr, v)
		if err != nil {
			return err
		}
	}
	return nil
}
