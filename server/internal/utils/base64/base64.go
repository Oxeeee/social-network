package base64encode

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func ToBase64(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	return encoded, nil
}

func FromBase64(base64Photo, username string) (string, error) {
	cleanBase64 := strings.Split(base64Photo, ",")

	decodedPhoto, err := base64.StdEncoding.DecodeString(cleanBase64[1])
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("avatars/userID-%s.png", username)
	err = os.WriteFile(fileName, decodedPhoto, 0664)
	if err != nil {
		return "", err
	}
	return fileName, err
}
