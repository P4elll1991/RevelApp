package controllers

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type JournalMapper struct {
	db *sql.DB
}

// метод получающий данные журнала

func (m JournalMapper) TakeJournal() ([]Event, error) {
	var err error
	m.db, err = InitDB() // ининциализация
	if err != nil {
		return nil, err
	}

	connStr := "SELECT * From journal"
	rows, err := m.db.Query(connStr) // запрос
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	journal := []Event{}

	for rows.Next() {
		p := Event{}
		err := rows.Scan(&p.Id, &p.Event, &p.BookId, &p.EmployeeId, &p.DateEvent, &p.IsbnJ, &p.BookNameJ, &p.NameJ, &p.CellnumberJ)
		if err != nil {
			return nil, err
		}

		journal = append(journal, p)
	}

	return journal, nil
}

// метод добавляющий нвовое событие в БД

func (m JournalMapper) AddEvent(event Event) error {
	var err error
	m.db, err = InitDB() // Инициализация БД
	if err != nil {
		return err
	}

	// Добавить елемент

	connStr := "insert into journal (event, bookid, isbn, BookName, employeeid, Name, Cellnumber, dateevent) values ( $1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = m.db.Exec(connStr, event.Event, event.BookId, event.IsbnJ, event.BookNameJ, event.EmployeeId, event.NameJ, event.CellnumberJ, event.DateEvent)

	if err != nil {
		return err
	}

	return nil
}
