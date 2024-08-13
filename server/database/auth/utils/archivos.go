package utils

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"

	"github.com/vincent-petithory/dataurl"
)

func isImage(mimeType string) bool {
	switch mimeType {
	case "image/jpeg", "image/png", "image/bmp", "image/webp":
		return true
	default:
		return false
	}
}

func SubirImagen(img64, prefix, idbol string) (string, error) {
	mydataurl, err := dataurl.DecodeString(img64)
	if err != nil {
		return "", err
	}

	if !isImage(mydataurl.MediaType.ContentType()) {
		return "", errors.New("el archivo debe ser una imagen")
	}

	sizeInBytes := len(mydataurl.Data)
	sizeInKB := float64(sizeInBytes) / 1024.0
	if sizeInKB > 1024 {
		return "", errors.New("la imagen no debe exceder 1 MB")
	}

	dir := "res"
	fileName := prefix + "-" + idbol + ".png"
	filePath := filepath.Join(dir, fileName)

	// Crear la carpeta si no existe
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(mydataurl.Data)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func GetImagen(url string) (string, error) {
	data, err := os.ReadFile(url)
	if err != nil {
		return "", err
	}

	// Codifica el archivo en base64
	base64Data := base64.StdEncoding.EncodeToString(data)

	// Construye el Data URI
	dataURI := "data:image/png;base64," + base64Data
	return dataURI, nil
}
