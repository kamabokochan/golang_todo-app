package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	// qmark style (?,?,?) 引数を渡している
	cmd := `insert into users (
        uuid,
        name,
        email,
        password,
        created_at) values (?, ?, ?, ?, ?)`

	// Dbは modelsパッケージ内のため、参照可能
	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
