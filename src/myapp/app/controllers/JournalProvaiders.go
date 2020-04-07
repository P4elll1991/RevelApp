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

func (pr JournalProvaider) GiveJournalPro() (journal []Event, err error) {

	journal, err = pr.mapper.TakeJournal()
	if err != nil {
		return nil, err
	}

	return journal, nil
}
