package main

import (
	"fmt"
	"log"
	"todo_app/src/app/models"
)

func main() {
	fmt.Println(models.Db)
	// controllers.StartMainServer()

	user, _ := models.GetUserByEmail("user1@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)
}
