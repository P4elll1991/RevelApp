package controllers

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type JournalMapper struct {
	db *sql.DB
}

func (m JournalMapper) TakeJournal() ([]Event, error) {
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer m.db.Close()

	connStr = "SELECT * From journal"
	rows, err := m.db.Query(connStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	journal := []Event{}

	for rows.Next() {
		p := Event{}
		err := rows.Scan(&p.Id, &p.Event, &p.BookId, &p.EmployeeId, &p.DateEvent, &p.Isbn, &p.BookName, &p.Name, &p.Cellnumber)
		if err != nil {
			return nil, err
		}

		journal = append(journal, p)
	}

	return journal, nil
}

func (m JournalMapper) AddEvent(event Event) error {
	var err error
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	m.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer m.db.Close()

	// Добавить елемент

	connStr = "insert into journal (event, bookid, isbn, BookName, employeeid, Name, Cellnumber, dateevent) values ( $1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = m.db.Exec(connStr, event.Event, event.BookId, event.Isbn, event.BookName, event.EmployeeId, event.Name, event.Cellnumber, event.DateEvent)

	if err != nil {
		return err
	}

	return nil
}
