package main

import (
	"fmt"
	"todo-app/app/controllers"
	"todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
