package main

import (
	"fmt"

	"github.com/RioTsukiji/ToDoApp/app/controllers"
	"github.com/RioTsukiji/ToDoApp/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
