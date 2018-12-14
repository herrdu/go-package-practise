package test

import (
	"io"
	"os"
	"strings"
)

func NewReader() error {
	r := strings.NewReader("some io.Reader stream to be read \n")

	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		return nil
	}
	defer file.Close()

	if _, err := io.Copy(file, r); err != nil {
		return nil
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return nil
	}
	return nil
}
