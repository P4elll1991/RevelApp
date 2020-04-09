package controllers

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

type Journal struct {
	*revel.Controller
	provaider JournalProvaider
}

// метод получающий данные из журнала

func (c Journal) Give() revel.Result {

	Journal, err := c.provaider.GiveJournalPro()
	if err != nil {
		fmt.Println(err)
	}
	return c.RenderJSON(Journal)
}

// методо доавляющий данные в журнал

func (c Journal) Add() revel.Result {

	var event Event
	c.Params.BindJSON(&event)

	err := c.provaider.AddEventPro(event)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render()
}
