package controllers

import (
	"time"

	_ "github.com/lib/pq"
)

type JournalProvaider struct {
	mapper JournalMapper
}

func (pr JournalProvaider) AddEventPro(event Event) error {

	event.DateEvent = time.Now()

	err := pr.mapper.AddEvent(event)
	if err != nil {
		return err
	}

	return nil
}

func (pr JournalProvaider) GiveJournalPro() (journal []EventPro, err error) {

	journalPro, err := pr.mapper.TakeJournal()
	if err != nil {
		return nil, err
	}

	for _, val := range journalPro {
		p := EventPro{}

		p.BookNameJ = val.BookName
		p.IsbnJ = val.Isbn
		p.NameJ = val.Name
		p.CellnumberJ = val.Cellnumber
		p.Id = val.Id
		p.Event = val.Event
		p.BookId = val.BookId
		p.EmployeeId = val.EmployeeId
		p.DateEvent = val.DateEvent.Format("2006-01-02")

		journal = append(journal, p)
	}

	return journal, nil
}
