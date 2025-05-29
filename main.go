package main

import (
	"arcade-website/components"
	"arcade-website/internal/database-init"
	"arcade-website/model"
	"arcade-website/pkg/database"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type PageData struct {
	Players []model.Player
}

func main() {
	db, err := database_init.DatabaseConnection()
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

	e.POST("/upload", database.UploadScore(db))

	// e.GET("/search/:option", database.SearchUser(db))

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
