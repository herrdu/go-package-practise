package test

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func ReadAt() error {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 16)

	buf := make([]byte, 6)
	if _, err := s.ReadAt(buf, 10); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	return nil
}
