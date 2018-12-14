package test

import (
	"io"
	"os"
	"strings"
)

func CopyBuffer() error {
	r1 := strings.NewReader("CopyBuffer is identical to Copy except that it stages through the provided buffer (if one is required) rather than allocating a temporary one. If buf is nil, one is allocated; otherwise if it has zero length, CopyBuffer panics.	\n")
	r2 := strings.NewReader("second reader \n")

	buf := make([]byte, 11)

	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		return err
	}
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		return err
	}

	return nil
}
