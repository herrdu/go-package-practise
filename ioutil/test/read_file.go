package test

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFile() {
	content, err := ioutil.ReadFile("../io/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}
