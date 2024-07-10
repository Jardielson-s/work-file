package utils

import (
	"errors"
	"os"
)

type File struct {
	FileName string
	Text     string
}

func (file File) CreateFile(input File) (string, error) {
	err := os.WriteFile(input.FileName, []byte(input.Text), 0644)

	if err != nil {
		return "", errors.New("Error for create file")
	}
	return "File created with success", nil
}
