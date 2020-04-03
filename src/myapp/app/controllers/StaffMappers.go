package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type StaffMapper struct {
	db *sql.DB
}

func (m StaffMapper) UpStaff(staff EmployeePro) error {
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

func (m StaffMapper) TakeStaff() ([]EmployeePro, []BookOfEmployee, error) {
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
	rows, err = m.db.Query(connStr)
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

func (m StaffMapper) AddStaff(staff EmployeePro) error {
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

func (m StaffMapper) StaffDeleteSome(s []int) error {
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

func (m StaffMapper) StaffDeleteOne(id int) error {
	// Открытие базы данных
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer m.db.Close()

	connStr = "delete from staff where id = $1"

	// Удаление из базы данных
	_, err = m.db.Exec(connStr, id)
	if err != nil {
		return err
	}

	return nil
}
