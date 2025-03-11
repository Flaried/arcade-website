package main

import (
	"arcade-website/components"
	"arcade-website/handlers"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type PageData struct {
	Players []handlers.Player
}

func main() {
	db, err := handlers.DatabaseConnection()
	if err != nil {
		log.Fatal("Failed to get database connection..")
	}
	defer db.Close()
	fmt.Println("Connected to database")

	e := echo.New()

	e.GET("/submit/:game_id", func(c echo.Context) error {
		gameID := c.Param("game_id")
		print(gameID)
		component := components.ScoreSubmission(gameID)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})
	e.POST("/upload", handlers.UploadScore(db))

	fmt.Println("Started Server on :6969")
	e.Logger.Fatal(e.Start("localhost:6969"))
}
