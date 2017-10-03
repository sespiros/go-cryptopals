package set2

import (
	"bytes"
)

// PKCSPadding implements PKCS padding for given data and key size
func PKCSPadding(data []byte, size int) []byte {
	l := size - len(data)%size

	return append(data, bytes.Repeat([]byte{byte(l)}, l)...)
}
