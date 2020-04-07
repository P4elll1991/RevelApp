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
}

type Auth struct {
	id       int
	username string
	password string
}

var auth = Auth{1, "admin", "pass"}

func (c App) Index() revel.Result {

	connStr := "user=postgres password=q dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	connStr = "SELECT * FROM auth"
	rows, err := db.Query(connStr)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	authList := []Auth{}

	for rows.Next() {
		p := Auth{}
		err := rows.Scan(&p.id, &p.username, &p.password)
		if err != nil {
			fmt.Println(err)
			continue
		}

		authList = append(authList, p)
	}
	check := false

	for _, val := range authList {
		if val.username == c.Session["username"] {
			check = true
		}
	}

	if check == true || c.Session["username"] == "???" {
		d := c.Request.GetHttpHeader("Authorization")
		d1 := strings.Split(d, ",")

		for _, val := range d1 {
			fmt.Println(val)
		}
		if len(d1) > 8 {
			resp := strings.Split(d1[4], "\"")
			username := strings.Split(d1[0], "\"")
			var password string
			check = false
			for _, val := range authList {
				if val.username == username[1] {
					check = true
					password = val.password
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
				fmt.Println(Resp)
				fmt.Println(HA1string)
				fmt.Println(ResponseStr)
				if Resp == resp[1] {
					c.Session["username"] = username[1]
					return c.Render()
				}
			}

		}
	}

	t := time.Now().Unix()
	rand := rand.Intn(999999999)
	nonceStr := fmt.Sprintf("%d:%d:privateKey", t, rand)
	nonce := md5.Sum([]byte(nonceStr))

	opaque := md5.Sum([]byte("RainbowSoft"))
	s := fmt.Sprintf("Digest realm=\"localhost:9000\",\n qop=\"auth,auth-int\",\n nonce=\"%x\",\n opaque=\"%x\"", nonce, opaque)

	c.Response.Out.Header().Add("WWW-Authenticate", s)
	c.Response.Status = 401
	c.Session["username"] = "???"
	return c.Render()

}

func (c App) Exite() revel.Result {

	c.Session["username"] = ""

	return c.Render()
}
