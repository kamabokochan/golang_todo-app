package main

import (
	"fmt"
	"todo_app/src/app/controllers"
	"todo_app/src/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
