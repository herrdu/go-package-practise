package test

import (
	"io"
	"log"
	"os"
	"strings"
)

func LimitReader() error {
	r := strings.NewReader("some io.Reader stream to be read\n")

	lr := io.LimitReader(r, 12)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
	return nil
}
