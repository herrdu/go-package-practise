package test

import (
	"io"
	"os"
)

func WriteString() error {
	io.WriteString(os.Stdout, "Hello World")
	return nil
}
