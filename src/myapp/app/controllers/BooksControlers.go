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

// Метод загружающий данные в таблицу с книгами

func (c Books) Give() revel.Result {
	allbooks, err := c.provaider.GiveBooksPro()
	if err != nil {
		fmt.Println(err)
	}
	return c.RenderJSON(allbooks)
}

// Метод добваляющий новую книгу

func (c Books) Add() revel.Result {
	var bookAdd Book
	c.Params.BindJSON(&bookAdd)

	err := c.provaider.AddBookPro(bookAdd)
	if err != nil {
		fmt.Println(err)
	}
	allbooks, err := c.provaider.GiveBooksPro()
	if err != nil {
		fmt.Println(err)
	}
	return c.RenderJSON(allbooks)
}

// Метод удаляющий книги

func (c Books) Delete() revel.Result {

	var Id []int // срез id книг

	c.Params.BindJSON(&Id)

	err := c.provaider.BookDeletePro(Id)
	if err != nil {
		fmt.Println(err)

	}

	allbooks, err := c.provaider.GiveBooksPro() // перезагрузка данных в таблице
	if err != nil {
		fmt.Println(err)
	}
	return c.RenderJSON(allbooks)
}

// метод обновляющий данные в таблице

func (c Books) Update() revel.Result {

	var bookUpdete Book
	c.Params.BindJSON(&bookUpdete)

	err := c.provaider.UpdateBookPro(bookUpdete) // обновление данных
	if err != nil {
		fmt.Println(err)
	}
	allbooks, err := c.provaider.GiveBooksPro() // перезагрузка таблицы с книгами
	if err != nil {
		fmt.Println(err)
	}
	Emp := StaffProvaider{}
	staff, err := Emp.GiveStaffPro() // перезагрузка данных с сотрудниками
	if err != nil {
		fmt.Println(err)
	}
	journalProvaider := JournalProvaider{}
	Journal, err := journalProvaider.GiveJournalPro() // перезагрузка журнала
	if err != nil {
		fmt.Println(err)
	}
	data := data{allbooks, staff, Journal} // объединенные данные книг, сотруднков и журнала
	return c.RenderJSON(data)
}
