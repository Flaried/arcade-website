package main

import (
	"fmt"
	"log"
	"lyles_arcade/handlers"

	"github.com/labstack/echo/v4"
)

func helloWorld(context echo.Context) error {
	return context.String(200, "Hello World!")
}

func main() {

	dbCon, err := handlers.DatabaseConnection()
	if err != nil {
		log.Fatal("Failed to get database connection..")
	}
	e := echo.New()
	e.GET("/", helloWorld)

	fmt.Println("Started Server..")
	e.Logger.Fatal(e.Start(":6969"))
}
