package controllers

import (
	"fmt"
)

type StaffProvaider struct {
	mapper StaffMapper
}

func (pr StaffProvaider) AddStaffPro(staff Employee) error {

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
		b := []Book{}
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

func (pr StaffProvaider) StaffDeletePro(staff []int) error {

	err := pr.mapper.StaffDelete(staff)
	if err != nil {
		return err
	}
	return nil
}

func (pr StaffProvaider) UpStaffPro(staff Employee) error {

	err := pr.mapper.UpStaff(staff)
	if err != nil {
		return err
	}
	return nil
}
