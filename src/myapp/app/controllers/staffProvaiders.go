package controllers

import (
	"fmt"
	"strconv"
)

type StaffProvaider struct {
	mapper StaffMapper
}

func (pr StaffProvaider) AddStaffPro(staff EmployeePro) error {

	err := pr.mapper.AddStaff(staff)
	if err != nil {
		return err
	}
	return nil
}

func (pr StaffProvaider) GiveStaffPro() (staff []Employee, err error) {

	staffPro, books, err := pr.mapper.TakeStaff()
	fmt.Println(books)
	if err != nil {
		return nil, err
	}
	for _, val := range staffPro {
		p := Employee{}
		b := []BookOfEmployee{}
		p.Id = val.Id
		p.Name = val.Name
		p.Department = val.Department
		p.Position = val.Position
		p.Cellnumber = val.Cellnumber

		for _, v := range books {
			if val.Id == v.Employeeid {
				b = append(b, v)
			}
		}

		p.Books = b
		staff = append(staff, p)
	}
	return staff, nil
}

func (pr StaffProvaider) StaffDeletePro(staff IdStaff) error {

	if staff.IdEmp != "" {
		Id, err := strconv.Atoi(staff.IdEmp)
		if err != nil {
			return err
		}
		err = pr.mapper.StaffDeleteOne(Id)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := pr.mapper.StaffDeleteSome(staff.IdStaff)
		if err != nil {
			return err
		}
		return nil
	}
}

func (pr StaffProvaider) UpStaffPro(staff EmployeePro) error {

	err := pr.mapper.UpStaff(staff)
	if err != nil {
		return err
	}
	return nil
}
