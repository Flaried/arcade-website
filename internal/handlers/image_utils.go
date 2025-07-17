package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func SavePicture(imageBytes []byte, path string) error {
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

func ValidatePicture(imageForm *multipart.FileHeader) ([]byte, error) {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	fileSrc, err := imageForm.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file %s", err.Error())
	}

	// Error doesn't matter because we are only reading
	defer fileSrc.Close()

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
