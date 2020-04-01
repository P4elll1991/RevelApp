package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/revel/revel"
)

type IdStaff struct {
	IdEmp   string
	IdStaff []int
}

func (c Staff) Delete() revel.Result {
	var IdArr IdStaff
	IdArr.IdEmp = c.Params.Query.Get("id")
	if IdArr.IdEmp != "" {
		err := StaffDeletePro(IdArr)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		c.Params.BindJSON(&IdArr.IdStaff)

		err := StaffDeletePro(IdArr)
		if err != nil {
			fmt.Println(err)
		}
	}

	return c.Render()
}

func StaffDeletePro(staff IdStaff) error {
	if staff.IdEmp != "" {
		Id, err := strconv.Atoi(staff.IdEmp)
		if err != nil {
			return err
		}
		err = StaffDelete1(Id)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := StaffDelete2(staff.IdStaff)
		if err != nil {
			return err
		}
		return nil
	}
}

func StaffDelete2(s []int) error {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	connStr = "delete from staff where id = $1"

	if err != nil {
		return err
	}
	defer db.Close()

	for _, v := range s {
		_, err = db.Exec(connStr, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func StaffDelete1(id int) error {
	// Открытие базы данных

	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	connStr = "delete from staff where id = $1"

	// Удаление из базы данных
	_, err = db.Exec(connStr, id)
	if err != nil {
		return err
	}

	return nil
}
