package test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Token() {
	const input = "1,2,3,4"
	scanner := bufio.NewScanner(strings.NewReader(input))

	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Println("advance", advance)
		fmt.Println("token", token)
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)

	for scanner.Scan() {
		fmt.Printf("scanner %q \n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
