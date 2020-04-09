package controllers

type AppPro struct {
	mapper AppMapper
}

// Провайдер получения данных с именами и паролями

func (p AppPro) GiveAuthDataPro() ([]Auth, error) {
	auth, err := p.mapper.GiveAuthData()
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// Провайдер добавления нового пользователя

func (p AppPro) AddAuthDataPro(user Auth) error {
	err := p.mapper.AddUserData(user)
	if err != nil {
		return err
	}
	return nil
}

// провайдер обновления текущего пользователя

func (p AppPro) UpdateAuthDataPro(user Auth, userSession interface{}) error {
	err := p.mapper.UpdateUserData(user, userSession)
	if err != nil {
		return err
	}
	return nil
}
