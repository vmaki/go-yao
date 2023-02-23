package helpers

import (
	"crypto/rand"
	"io"
)

func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, length)
	_, _ = io.ReadAtLeast(rand.Reader, b, length)
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}
