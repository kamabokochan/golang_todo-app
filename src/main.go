package main

import (
	"fmt"
	"todo_app/src/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)
	// sqlite3 webapp.sql
	// sqlite> .table
	// usersと表示される

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "testtest"
	fmt.Println(u)

	u.CreateUser()
	// &{0  test test@example.com testtest 0001-01-01 00:00:00 +0000 UTC}
	// SELECT * FROM users;
	// id, uuid, name, email, password, createdAt
	// 1|44242a8c-8ba1-11ee-8768-38f9d35db3a1|test|test@example.com|51abb9636078defbf888d8457a7c76f85c8f114c|2023-11-25 23:45:24.42818+09:00
}
