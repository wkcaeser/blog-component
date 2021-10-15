package user

import (
	"blog-component/db"
	"log"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Icon     string `json:"icon"`
}

type UserCrud interface {
	CreateNewUser() int64
}

func (user *User) CreateNewUser() int64 {

	rs, e := db.DbPool.Exec("insert into user(create_time, modify_time, username, nick, email, password, icon, status)"+
		" value(now(), now(), ?, ?, ?, ?, ?, 1)", user.Username, user.Nick, user.Email, user.Password, user.Icon)

	if e != nil {
		log.Panicln("user insert error", e.Error())
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}
