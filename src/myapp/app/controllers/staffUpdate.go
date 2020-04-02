package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

func (c Staff) Update() revel.Result {
	EmployeeProvader := EmployeePro{}
	var empUpdate EmployeePro
	c.Params.BindJSON(&empUpdate)

	err := EmployeeProvader.UpStaffPro(empUpdate)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(empUpdate)
	return c.Render()
}

func (EmployeePro) UpStaffPro(staff EmployeePro) error {
	EmployeeMapper := Employee{}
	err := EmployeeMapper.UpStaff(staff)
	if err != nil {
		return err
	}
	return nil
}

func (Employee) UpStaff(staff EmployeePro) error {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Добавить елемент

	connStr = "update staff set name = $1, department = $2, position = $3, cellnumber = $4 where id = $5"
	_, err = db.Exec(connStr, staff.Name, staff.Department, staff.Position, staff.Cellnumber, staff.Id)

	if err != nil {
		return err
	}

	return nil
}
