package handlers

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
)

func validatePicture(imageForm *multipart.FileHeader) ([]byte, error) {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
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
