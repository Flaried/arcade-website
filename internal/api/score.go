package api

import (
	"arcade-website/internal/handlers"
	"arcade-website/internal/model"
	"database/sql"
	"fmt"
	"log"
	// "os"

	"github.com/labstack/echo/v4"
)

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

		imageBytes, err := handlers.ValidatePicture(userFile)
		if err != nil {
			c.Response().Header().Set("HX-Reswap", "innerHTML")
			fmt.Println(err.Error())
			return c.HTML(200, `<div class="text-red-500">Error Sanaizing file</div>`)
		}

		savedPath := fmt.Sprintf("scores/%s_%s_%s_%s.jpeg", Submission.GameID, Submission.Username, Submission.Score, "pending")

		err = handlers.SavePicture(imageBytes, savedPath)
		if err != nil {
			return c.HTML(200, err.Error())
		}

		fmt.Println("Submission", Submission, Submission.Username)
		var userID int
		err = db.QueryRow("SELECT user_id FROM users WHERE username = $1", Submission.Username).Scan(&userID)

		if err != nil {
			if err == sql.ErrNoRows {
				// User doesn't exist → prompt to create new user
				// TODO: IMPLEMENT CREATE USER PAGE OR API
				c.Response().Header().Set("HX-Reswap", "innerHTML")
				return c.HTML(200, fmt.Sprintf(`
            <div class="text-yellow-500">
                User "%s" not found. Do you want to create a new user?
                <button hx-post="/create-user" hx-vals='{"username":"%s","initials":"%s"}'>Yes</button>
            </div>
        `, Submission.Username, Submission.Username, Submission.Initials))
			} else {
				// Some other DB error
				fmt.Println("DB error:", err)
				return c.HTML(200, `<div class="text-red-500">Error querying database</div>`)
			}
		}
		// NOTE: THIS JUST MKAES A USER, DOESNT SAVE A SCORE.
		// _, err = db.Exec("INSERT INTO users (username, initials) VALUES ($1, $2)", Submission.Username, Submission.Initials)
		// if err != nil {
		// 	err = os.Remove(savedPath)
		// 	if err != nil {
		// 		fmt.Printf("Couldn't remove %s\n", savedPath)
		// 	}
		//
		// 	c.Response().Header().Set("HX-Reswap", "innerHTML")
		// 	return c.HTML(200, `<div class="text-red-500">Error Writing to Database File</div>`)
		// }

		return c.HTML(200, `<div class="text-green-500">Sucessfully Published Score</div>`)
	}
}
