package main

import (
	"arcade-website/internal/api"
	"arcade-website/internal/database"
	"arcade-website/internal/templates"
	"arcade-website/internal/templates/submit"

	"github.com/a-h/templ"

	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func renderTempl(c echo.Context, component templ.Component) error {
	c.Response().Header().Set("Content-Type", "text/html")
	c.Response().WriteHeader(c.Response().Status)
	return component.Render(c.Request().Context(), c.Response().Writer)
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
		content := submit.ScoreSubmission(gameID)

		page := templates.Base("Score Submission", content)
		return renderTempl(c, page)
	})

	// APIs
	e.POST("/api/upload", api.PostScore(db))

	// e.GET("/search/:option", database.SearchUser(db))
	e.GET("/api/set-username", func(c echo.Context) error {
		username := c.QueryParam("username")
		if username == "" {
			fmt.Println("Is initials")
		}
		component := submit.UsernameInput(username)
		return renderTempl(c, component)

		// return component.Render(c.Request().Context(), c.Response().Writer)
	})

	fmt.Println("Started Server on :6969")
	e.Logger.Fatal(e.Start("localhost:6969"))
}
