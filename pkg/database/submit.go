package database

import (
	"bytes"
	"database/sql"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func sanatizePicture(imageForm *multipart.FileHeader) ([]byte, error) {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}

	fileSrc, err := imageForm.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file %s", err.Error())
	}

	defer fileSrc.Close() // Always close the file

	fileBytes, err := io.ReadAll(fileSrc)
	if err != nil {
		return nil, fmt.Errorf("cannot read file bytes %s", err.Error())
	}

	detectedType := http.DetectContentType(fileBytes)
	if !allowedTypes[detectedType] {
		return nil, fmt.Errorf("invalid file types: %s", detectedType)
	}

	imageDecoded, err := jpeg.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		return nil, fmt.Errorf("cannot decode Image... %s", err.Error())
	}

	var buf bytes.Buffer

	err = jpeg.Encode(&buf, imageDecoded, &jpeg.Options{Quality: 25})
	if err != nil {
		return nil, fmt.Errorf("cannot Encode Image... %s", err.Error())
	}

	return buf.Bytes(), nil
}

type FormSubmission struct {
	GameID   string
	Score    string
	Username string
	Initials string
}

func UploadScore(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//var serverMessages []string
		var Submission FormSubmission
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

// func SearchUser(db *sql.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var userForm components.UserSearchForm
// 		users, err := GetAllUsers(db)
// 		fmt.Println(users)
// 		if err != nil {
// 			userForm.GeneralErrors = append(userForm.GeneralErrors, "Database error: "+err.Error())
//
// 			return components.UserResults(userForm).Render(c.Request().Context(), c.Response().Writer)
// 		}
//
// 		optionQuery := c.Param("option")
// 		query := c.QueryParam("inputData")
// 		fmt.Println(optionQuery, "query option")
// 		if optionQuery == "initials" {
// 			for _, user := range users {
// 				if strings.Contains(strings.ToLower(user.Initials), strings.ToLower(query)) {
// 					userForm.Fields = append(userForm.Fields, user.Username)
// 				}
// 			}
//
// 		} else {
// 			for _, user := range users {
// 				if strings.Contains(strings.ToLower(user.Username), strings.ToLower(query)) {
// 					userForm.Fields = append(userForm.Fields, user.Username)
// 				}
// 			}
// 		}
//
// 		return components.UserResults(userForm).Render(c.Request().Context(), c.Response().Writer)
// 	}
// }
