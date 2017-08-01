package set2

import (
	"bytes"
)

func PKCSpadding(data []byte, size int) []byte {
	l := size - len(data)%size

	return append(data, bytes.Repeat([]byte{byte(l)}, l)...)
}
