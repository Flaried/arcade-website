package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type PageData struct {
	Players []handlers.Player
}

func main() {
	db, err := database.DatabaseConnection()
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

	e.GET("/search/:option", handlers.SearchUser(db))

	e.GET("/set-username", func(c echo.Context) error {
		username := c.QueryParam("username")
		if username == "" {
			fmt.Println("Is initials")
		}
		component := components.UsernameInput(username)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	fmt.Println("Started Server on :6969")
	e.Logger.Fatal(e.Start("localhost:6969"))
}
