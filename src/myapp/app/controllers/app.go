package controllers

import (
	"crypto/md5"
	"database/sql"
	"fmt"

	"math/rand"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
	provaider AppPro
}

// Структура аутентификации

type Auth struct {
	id       int
	Username string
	Password string
}

var DB *sql.DB

// Функция инициализации БД

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres password=q dbname=library sslmode=disable"
	var err error
	if DB == nil {
		DB, err = sql.Open("postgres", connStr)
	}

	return DB, err
}

// Главный метод. Производит аутентификацию и загружает index.html

func (c App) Index() revel.Result {
	check := false // Флаг для проверки существования имени пользователя в сессии в БД
	authList, err := c.provaider.GiveAuthDataPro()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range authList {
			if val.Username == c.Session["username"] {
				check = true
			}
		}
	}

	// Если в сессии хранится существующее имя пользователя или "???" (Передается в сессию при отправки ответа 401)то осуществляется проверка данных в заголовке запроса

	if check == true || c.Session["username"] == "???" {
		d := c.Request.GetHttpHeader("Authorization")
		d1 := strings.Split(d, ",")

		if len(d1) > 8 {
			resp := strings.Split(d1[4], "\"")
			username := strings.Split(d1[0], "\"")
			var password string
			check = false
			for _, val := range authList {
				if val.Username == username[1] {
					check = true
					password = val.Password
				}
			}
			if check == true {
				cnonce := strings.Split(d1[8], "\"")
				nonce := strings.Split(d1[2], "\"")
				nc := strings.Split(d1[7], "=")
				qop := strings.Split(d1[6], "=")
				HA1string := fmt.Sprintf("%s:localhost:9000:%s", username[1], password)
				HA1 := md5.Sum([]byte(HA1string))
				HA2 := md5.Sum([]byte("GET:/"))
				ResponseStr := fmt.Sprintf("%x:%s:%s:%s:%s:%x", HA1, nonce[1], nc[1], cnonce[1], qop[1], HA2)
				Responser := md5.Sum([]byte(ResponseStr))
				Resp := fmt.Sprintf("%x", Responser)

				if Resp == resp[1] {
					c.Session["username"] = username[1]
					c.ViewArgs["userVar"] = username[1]
					return c.Render()
				}
			}
		}
	}

	// Формирование данных для дайджест протокола
	t := time.Now().Unix()       // Время отправки ответа
	rand := rand.Intn(999999999) // Случайное число
	nonceStr := fmt.Sprintf("%d:%d:privateKey", t, rand)
	nonce := md5.Sum([]byte(nonceStr)) // Унтикальная переменная для данного пользователя

	opaque := md5.Sum([]byte("RainbowSoft"))                                                                                     // Переменная известная в данном сервере
	s := fmt.Sprintf("Digest realm=\"localhost:9000\",\n qop=\"auth,auth-int\",\n nonce=\"%x\",\n opaque=\"%x\"", nonce, opaque) // заголовок дайджест протокола

	c.Response.Out.Header().Add("WWW-Authenticate", s)
	c.Response.Status = 401
	c.Session["username"] = "???" // В сессиию передается информация что осуществлется аутентификация
	return c.Render()

}

func (c App) Exite() revel.Result {

	c.Session["username"] = "" // Выход из приложения

	return c.Render()
}

// Метод добавление нового пользователя

func (c App) Add() revel.Result {

	var userDate Auth
	c.Params.BindJSON(&userDate)

	check, err := Check(userDate.Username) // Проверка сущетсования имени пользователя
	if err != nil {
		fmt.Println(err)
		mess := "2"
		return c.RenderText(mess)
	}

	if !check { // Если имя пользователя уже существует то отправляется сообщение "2" что сигнализирует клиента вывеси сообщение об ошибке
		mess := "2"
		return c.RenderText(mess)
	}

	err = c.provaider.AddAuthDataPro(userDate) // Добавление нового пользователя
	if err != nil {
		fmt.Println(err)
	}

	mess := "1"
	return c.RenderText(mess)
}

// Метод обновления данных пользователя

func (c App) Update() revel.Result {

	var userDate Auth
	c.Params.BindJSON(&userDate)
	userSession := c.Session["username"]

	check, err := Check(userDate.Username) // Проверка сущетсования имени пользователя
	if err != nil {
		fmt.Println(err)
		mess := "2"
		return c.RenderText(mess)
	}

	if !check { // Если имя пользователя уже существует то отправляется сообщение "2" что сигнализирует клиента вывеси сообщение об ошибке
		mess := "2"
		return c.RenderText(mess)
	}

	err = c.provaider.UpdateAuthDataPro(userDate, userSession) // Обновление данных пользователя
	if err != nil {
		fmt.Println(err)
	}

	mess := "1"
	return c.RenderText(mess)
}

// Функция проверки существования имени пользователя в БД

func Check(username string) (bool, error) {

	db, err := InitDB() // Инициализация БД
	if err != nil {
		fmt.Println(err)
	}

	connStr := "select * from auth where username =$1"
	rows, err := db.Query(connStr, username) // Нахождение строки с нужным именнем
	if err != nil {
		fmt.Println(err)
		return true, err
	}
	defer rows.Close()

	usersDate := []Auth{}

	for rows.Next() { //перевод найденных строк в структуру аутентификации
		p := Auth{}
		err := rows.Scan(&p.id, &p.Username, &p.Password)
		if err != nil {
			fmt.Println(err)
			continue
		}

		usersDate = append(usersDate, p)
	}

	if len(usersDate) != 0 { // если длина среза больше 0 (то есть строка найдена) то отпарвляется ложь
		return false, nil
	}
	// иначе истина
	return true, nil
}
