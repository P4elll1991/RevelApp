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

func (c Journal) Give() revel.Result {

	Journal, err := c.provaider.GiveJournalPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(Journal)
}

func (c Journal) Add() revel.Result {

	var event Event
	c.Params.BindJSON(&event)

	err := c.provaider.AddEventPro(event)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render()
}
