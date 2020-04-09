package controllers

type StaffProvaider struct {
	mapper StaffMapper
}

// метод добавляющий нового сотрудника

func (pr StaffProvaider) AddStaffPro(staff Employee) error {

	err := pr.mapper.AddStaff(staff)
	if err != nil {
		return err
	}
	return nil
}

// метод получающий данные о сотруднках

func (pr StaffProvaider) GiveStaffPro() (staff []Employee, err error) {

	staffPro, books, err := pr.mapper.TakeStaff() // получение срезов сотрудников и книг

	if err != nil {
		return nil, err
	}
	for _, val := range staffPro { // Добавление данных книг к срезу сотрудников
		p := Employee{}
		b := []Book{}
		p.Id = val.Id
		p.Name = val.Name
		p.Department = val.Department
		p.Position = val.Position
		p.Cellnumber = val.Cellnumber

		for _, v := range books { // формирование среза книг имеющихся у данного сотрудника
			if val.Id == v.Employeeid {
				b = append(b, v)
			}
		}

		p.Books = b // добавление полученного среза книг к сотруднику
		staff = append(staff, p)
	}
	return staff, nil
}

// метод удаление сотрудника

func (pr StaffProvaider) StaffDeletePro(staff []int) error {

	err := pr.mapper.StaffDelete(staff)
	if err != nil {
		return err
	}
	return nil
}

// метод обновления сотрудника

func (pr StaffProvaider) UpStaffPro(staff Employee) error {

	err := pr.mapper.UpStaff(staff)
	if err != nil {
		return err
	}
	return nil
}
