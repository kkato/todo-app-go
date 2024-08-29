package main

import (
	"fmt"

	"github.com/kkato/todo-app/app/controllers"
	"github.com/kkato/todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
