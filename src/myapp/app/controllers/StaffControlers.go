package controllers

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/revel/revel"
)

type Staff struct {
	*revel.Controller
	provaider StaffProvaider
}

// метод получающий даннные сотрудников

func (c Staff) Give() revel.Result {

	staff, err := c.provaider.GiveStaffPro()
	if err != nil {
		fmt.Println(err)
	}
	return c.RenderJSON(staff)
}

// метод добавляющий сотрудника

func (c Staff) Add() revel.Result {

	var empAdd Employee
	c.Params.BindJSON(&empAdd)

	err := c.provaider.AddStaffPro(empAdd)
	if err != nil {
		fmt.Println(err)
	}

	b := Books{}
	allbooks, err := b.provaider.GiveBooksPro() // перезагрузка таблицы книг
	if err != nil {
		fmt.Println(err)
	}
	staff, err := c.provaider.GiveStaffPro() // перезрузка таблицы сотрудников
	if err != nil {
		fmt.Println(err)
	}

	data := data{allbooks, staff, nil} // общие данные книг и сотрудников
	return c.RenderJSON(data)
}

// метод удаляющий сотрудников

func (c Staff) Delete() revel.Result {

	var Id []int

	c.Params.BindJSON(&Id)

	err := c.provaider.StaffDeletePro(Id)
	if err != nil {
		fmt.Println(err)
	}

	staff, err := c.provaider.GiveStaffPro() // перезрузка таблицы сотрудников
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(staff)
}

// метод обновляющий сотрудника

func (c Staff) Update() revel.Result {

	var empUpdate Employee
	c.Params.BindJSON(&empUpdate)

	err := c.provaider.UpStaffPro(empUpdate)
	if err != nil {
		fmt.Println(err)
	}
	b := Books{}
	allbooks, err := b.provaider.GiveBooksPro() // перезагрузка таблицы книг
	if err != nil {
		fmt.Println(err)
	}
	staff, err := c.provaider.GiveStaffPro() // перезрузка таблицы сотрудников
	if err != nil {
		fmt.Println(err)
	}

	data := data{allbooks, staff, nil} // общие данные книг и сотрудников
	return c.RenderJSON(data)
}
