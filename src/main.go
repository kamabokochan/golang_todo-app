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
}
