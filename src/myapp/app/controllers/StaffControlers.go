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

func (c Staff) Give() revel.Result {

	staff, err := c.provaider.GiveStaffPro()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(staff)
}

func (c Staff) Add() revel.Result {

	var empAdd Employee
	c.Params.BindJSON(&empAdd)

	err := c.provaider.AddStaffPro(empAdd)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render()
}

func (c Staff) Delete() revel.Result {

	var Id []int

	c.Params.BindJSON(&Id)

	err := c.provaider.StaffDeletePro(Id)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}

func (c Staff) Update() revel.Result {

	var empUpdate Employee
	c.Params.BindJSON(&empUpdate)

	err := c.provaider.UpStaffPro(empUpdate)
	if err != nil {
		fmt.Println(err)
	}
	b := Books{}
	allbooks, err := b.provaider.GiveBooksPro()
	if err != nil {
		panic(err)
	}
	staff, err := c.provaider.GiveStaffPro()
	if err != nil {
		panic(err)
	}

	data := data{allbooks, staff, nil}
	return c.RenderJSON(data)
}
