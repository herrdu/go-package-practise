package test

import (
	"fmt"
	"io/ioutil"
	"log"
)

func TempFile() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmpfile.Name())

	// defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TempFileWithSuffix() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tmpfile.Name())

	// defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
