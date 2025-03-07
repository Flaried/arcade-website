package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UploadScore() echo.HandlerFunc {
	return func(c echo.Context) error {
		var serverMessages []string

		// Set maximum bytes of reader
		c.Request().Body = http.MaxBytesReader(c.Response().Writer, c.Request().Body, 30*1024*1024)

		fmt.Println(c.FormValue("game_id"))
		fmt.Println(c.FormValue("score"))
		fmt.Println(c.FormValue("username"))
		fmt.Println(c.FormValue("initials"))
		fmt.Println(serverMessages)

		file, err := c.FormFile("photo")
		if err != nil {
			fmt.Println("Photo too large")
			return c.HTML(400, `<div class="text-red-500">Missing file upload</div>`)
		}
		fmt.Println(file.Size)
		return c.NoContent(200)
	}
}
