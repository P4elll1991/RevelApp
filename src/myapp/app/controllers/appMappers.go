package controllers

import "database/sql"

type AppMapper struct {
	db *sql.DB
}

// метод получения из БД данных с именами и паролями

func (m AppMapper) GiveAuthData() ([]Auth, error) {
	var err error
	m.db, err = InitDB() // Инициализация БД
	if err != nil {
		return nil, err
	}

	connStr := "SELECT * FROM auth"
	rows, err := m.db.Query(connStr) // Запрос в БД
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authList := []Auth{}

	for rows.Next() { // Перевод строк в срез структур
		p := Auth{}
		err := rows.Scan(&p.id, &p.Username, &p.Password)
		if err != nil {
			return nil, err
		}

		authList = append(authList, p)
	}

	return authList, nil
}

// метод добавления нового пользователя

func (m AppMapper) AddUserData(user Auth) error {
	var err error
	m.db, err = InitDB() // Инициализация БД
	if err != nil {
		return err
	}

	connStr := "insert into auth (username, password) values ($1, $2)"
	rows, err := m.db.Query(connStr, user.Username, user.Password) // Запрос
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

//метод обновления текущего пользователя

func (m AppMapper) UpdateUserData(user Auth, userSession interface{}) error {
	var err error
	m.db, err = InitDB() //инициализация БД
	if err != nil {
		return err
	}

	connStr := "update auth set username = $1, password = $2 where username = $3"
	rows, err := m.db.Query(connStr, user.Username, user.Password, userSession) //Запрос
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
