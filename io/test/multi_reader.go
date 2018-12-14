package test

import (
	"io"
	"os"
	"strings"
)

func MultiReader() error {

	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil
}
