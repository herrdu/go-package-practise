package test

import (
	"io"
	"os"
	"strings"
)

func CopyN() error {

	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := io.CopyN(os.Stdout, r, 6); err != nil {
		return err
	}
	return nil
}
