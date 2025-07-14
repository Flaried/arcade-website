package main

import (
	"arcade-website/internal/database"
	"arcade-website/internal/handlers"
	"arcade-website/internal/templates"

	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

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
		component := templates.ScoreSubmission(gameID)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.POST("/upload", handlers.UploadScore(db))

	// e.GET("/search/:option", database.SearchUser(db))

	e.GET("/set-username", func(c echo.Context) error {
		username := c.QueryParam("username")
		if username == "" {
			fmt.Println("Is initials")
		}
		component := templates.UsernameInput(username)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	fmt.Println("Started Server on :6969")
	e.Logger.Fatal(e.Start("localhost:6969"))
}
