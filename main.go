package main

import (
	"github.com/kappaggaeru/seminaryGo/controller"
	"github.com/kappaggaeru/seminaryGo/model/database"
	"github.com/kappaggaeru/seminaryGo/service"
)

func main() {
	db := database.NewDatabase("database.db")
	db.CreateSchema()

	service := service.NewInstance(db)

	http := controller.NewHTTPController(&service)
	http.Run()
}