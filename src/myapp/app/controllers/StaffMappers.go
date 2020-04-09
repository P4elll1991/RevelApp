package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type StaffMapper struct {
	db *sql.DB
}

// Обновление данных о сотрудниках

func (m StaffMapper) UpStaff(staff Employee) error {
	var err error
	m.db, err = InitDB() // инициализация БД
	if err != nil {
		return err
	}

	// запрос

	connStr := "update staff set name = $1, department = $2, position = $3, cellnumber = $4 where id = $5"
	_, err = m.db.Exec(connStr, staff.Name, staff.Department, staff.Position, staff.Cellnumber, staff.Id)

	if err != nil {
		return err
	}

	return nil
}

// метод получающий данные о сотрудниках

func (m StaffMapper) TakeStaff() ([]Employee, []Book, error) {
	var err error
	m.db, err = InitDB() // инициализация
	if err != nil {
		return nil, nil, err
	}

	connStr := "SELECT * FROM staff 	WHERE id != 1"
	rows, err := m.db.Query(connStr) // запрос на получение сотрудников
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	staff := []Employee{}

	for rows.Next() { // формирование среза структур сотрудников
		p := Employee{}
		err := rows.Scan(&p.Id, &p.Name, &p.Department, &p.Position, &p.Cellnumber)
		if err != nil {
			fmt.Println(err)
			continue
		}

		staff = append(staff, p)
	}

	connStr = "SELECT id, Isbn, BookName, Employeeid, Datestart From books"
	rows, err = m.db.Query(connStr) // запрос на получение данных книг
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	books := []Book{}

	for rows.Next() { // формирование среза книг
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

// метод добавление сотрудника

func (m StaffMapper) AddStaff(staff Employee) error {
	var err error
	m.db, err = InitDB() // инициализация БД
	if err != nil {
		return err
	}

	// Добавить елемент

	connStr := "insert into staff (name, department, position, cellnumber) values ( $1, $2, $3, $4)"
	_, err = m.db.Exec(connStr, staff.Name, staff.Department, staff.Position, staff.Cellnumber)

	if err != nil {
		return err
	}

	return nil
}

// Метод удаление сотрудника

func (m StaffMapper) StaffDelete(s []int) error {
	var err error
	m.db, err = InitDB() // Инициализация
	connStr := "delete from staff where id = $1"

	if err != nil {
		return err
	}

	for _, v := range s { // Перебор среза id
		_, err = m.db.Exec(connStr, v) // удаление елемента
		if err != nil {
			return err
		}
	}
	return nil
}
