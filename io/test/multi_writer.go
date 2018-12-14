package test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func MultiWriter() error {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf1, buf2 bytes.Buffer

	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
	return nil
}
