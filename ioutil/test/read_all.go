package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ReadAll() error {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

	return nil
}
