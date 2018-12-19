package test

import (
	"bufio"
	"fmt"
	"os"
)

func Write() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush()
}
