package main

import (
	"arcade-website/handlers"
	"arcade-website/templates"
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

	fmt.Println("Connected to server..")

	players, err := handlers.FetchUsers(db)
	if err != nil {
		log.Fatalf("Error fetching users: %s", err.Error())
	}
	fmt.Println(players)

	e := echo.New()
	e.Renderer = templates.NewTemplate()

	e.GET("/", func(c echo.Context) error {
		data := PageData{Players: players}
		return c.Render(200, "index", data)
	})

	e.GET("/submit/:game_id", func(c echo.Context) error {
		gameID := c.Param("game_id")
		return c.Render(200, "submit-index", map[string]any{
			"gameID": gameID,
		})
	})

	//e.Use(middleware.BodyLimit("10M"))
	e.POST("/upload", handlers.UploadScore())
	e.POST("/users", func(c echo.Context) error {
		fmt.Println(players)
		data := PageData{Players: players}
		return c.Render(200, "player-block", data)
	})

	fmt.Println("Started Server..")
	e.Logger.Fatal(e.Start("localhost:6969"))
}
