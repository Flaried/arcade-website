package handlers

import (
	"arcade-website/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

// TODO: return errors and html more gracefully
func PostScore(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//var serverMessages []string
		var Submission model.FormSubmission
		// Set maximum bytes of reader
		// c.Request().Body = http.MaxBytesReader(c.Response().Writer, c.Request().Body, 30*1024*1024)
		//
		Submission.Score = c.FormValue("score")
		Submission.Username = c.FormValue("username")
		Submission.Initials = c.FormValue("initials")
		Submission.GameID = c.FormValue("game_id")

		//TODO: EXIT/sanatize all form values

		/// User inputed bytes.
		userFile, err := c.FormFile("photo")
		if err != nil {
			log.Println(err.Error())
			// TODO: RETURN ERROR
			c.Response().Header().Set("HX-Reswap", "innerHTML")
			return c.HTML(200, `<div class="text-red-500">Missing file upload</div>`)
		}

		imageBytes, err := validatePicture(userFile)
		if err != nil {
			c.Response().Header().Set("HX-Reswap", "innerHTML")
			fmt.Println(err.Error())
			return c.HTML(200, `<div class="text-red-500">Error Sanaizing file</div>`)
		}

		savedPath := fmt.Sprintf("scores/%s_%s_%s_%s.jpeg", Submission.GameID, Submission.Username, Submission.Score, "pending")

		err = savePicture(imageBytes, savedPath)
		if err != nil {
			return c.HTML(200, err.Error())
		}

		fmt.Println("Submission", Submission, Submission.Username)
		_, err = db.Exec("INSERT INTO users (username, initials) VALUES ($1, $2)", Submission.Username, Submission.Initials)
		if err != nil {
			err = os.Remove(savedPath)
			if err != nil {
				fmt.Printf("Couldn't remove %s\n", savedPath)
			}

			c.Response().Header().Set("HX-Reswap", "innerHTML")
			return c.HTML(200, `<div class="text-red-500">Error Writing to Database File</div>`)
		}

		return c.HTML(200, `<div class="text-green-500">Sucessfully Published Score</div>`)
	}
}

func savePicture(imageBytes []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Println(err.Error())
		return errors.New(`<div class="text-red-500">Error Downloading file</div>`)
	}

	_, err = file.Write(imageBytes)
	if err != nil {
		log.Println(err.Error())
		// c.Response().Header().Set("HX-Reswap", "innerHTML")
		return errors.New(`<div class="text-red-500">Error Creating Image file</div>`)
	}

	err = file.Close()
	if err != nil {
		return errors.New(`<div class="text-red-500">Error closing file</div>`)
	}

	return nil
}
