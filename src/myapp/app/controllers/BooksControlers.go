package controllers

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

type Books struct {
	*revel.Controller
	provaider BookProvaider
}

func (c Books) Give() revel.Result {
	allbooks, err := c.provaider.GiveBooksPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(allbooks)
}
func (c Books) Add() revel.Result {
	var bookAdd Book
	c.Params.BindJSON(&bookAdd)

	err := c.provaider.AddBookPro(bookAdd)
	if err != nil {
		fmt.Println(err)
	}
	allbooks, err := c.provaider.GiveBooksPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(allbooks)
}

func (c Books) Delete() revel.Result {

	var Id []int

	c.Params.BindJSON(&Id)

	err := c.provaider.BookDeletePro(Id)
	if err != nil {
		fmt.Println(err)

	}

	allbooks, err := c.provaider.GiveBooksPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(allbooks)
}

func (c Books) Update() revel.Result {

	var bookUpdete Book
	c.Params.BindJSON(&bookUpdete)

	err := c.provaider.UpdateBookPro(bookUpdete)
	if err != nil {
		fmt.Println(err)
	}
	allbooks, err := c.provaider.GiveBooksPro()
	if err != nil {
		panic(err)
	}
	Emp := StaffProvaider{}
	staff, err := Emp.GiveStaffPro()
	if err != nil {
		panic(err)
	}
	journalProvaider := JournalProvaider{}
	Journal, err := journalProvaider.GiveJournalPro()
	if err != nil {
		panic(err)
	}
	data := data{allbooks, staff, Journal}
	return c.RenderJSON(data)
}
