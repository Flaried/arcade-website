package handlers

import (
	"arcade-website/internal/model"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func PostScore(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//var serverMessages []string
		var Submission model.FormSubmission
		// Set maximum bytes of reader
		// c.Request().Body = http.MaxBytesReader(c.Response().Writer, c.Request().Body, 30*1024*1024)
		//
		Submission.Score = c.FormValue("score")

		fmt.Println(c.FormValue)
		Submission.Username = c.FormValue("username")
		Submission.Initials = c.FormValue("initials")
		Submission.GameID = c.FormValue("game_id")

		//TODO: EXIT/sanatize all form values

		file, err := c.FormFile("photo")
		if err != nil {
			log.Println(err.Error())
			// TODO: RETURN ERROR
			return c.HTML(400, `<div class="text-red-500">Missing file upload</div>`)
		}

		imageBytes, err := sanatizePicture(file)
		if err != nil {
			fmt.Println(err.Error())
			return c.HTML(400, `<div class="text-red-500">Error Sanaizing file</div>`)
		}

		savedPath := fmt.Sprintf("scores/%s_%s_%s_%s.jpeg", Submission.GameID, Submission.Username, Submission.Score, "pending")
		f, err := os.Create(savedPath)
		if err != nil {
			log.Println(err.Error())
			return c.HTML(400, `<div class="text-red-500">Error Downloading file</div>`)
		}
		defer f.Close()

		_, err = f.Write(imageBytes)
		if err != nil {
			log.Println(err.Error())
			return c.HTML(400, `<div class="text-red-500">Error Creating Image file</div>`)
		}

		fmt.Println("Submission", Submission, Submission.Username)
		_, err = db.Exec("INSERT INTO users (username, initials) VALUES ($1, $2)", Submission.Username, Submission.Initials)
		if err != nil {
			return c.HTML(400, `<div class="text-red-500">Error Writing to Databases File</div>`)
		}

		return c.NoContent(200)
	}
}
