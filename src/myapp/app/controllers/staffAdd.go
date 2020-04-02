package controllers

import (
	"database/sql"
	"fmt"

	"github.com/revel/revel"
)

func (c Staff) Add() revel.Result {
	EmployeeProvader := EmployeePro{}
	var empAdd EmployeePro
	c.Params.BindJSON(&empAdd)

	err := EmployeeProvader.AddStaffPro(empAdd)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render()
}

func (EmployeePro) AddStaffPro(staff EmployeePro) error {
	EmployeeMapper := Employee{}
	err := EmployeeMapper.AddStaff(staff)
	if err != nil {
		return err
	}
	return nil
}

func (Employee) AddStaff(staff EmployeePro) error {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Добавить елемент

	connStr = "insert into staff (name, department, position, cellnumber) values ( $1, $2, $3, $4)"
	_, err = db.Exec(connStr, staff.Name, staff.Department, staff.Position, staff.Cellnumber)

	if err != nil {
		return err
	}

	return nil
}
