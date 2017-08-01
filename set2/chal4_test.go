package set2

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/sespiros/go-cryptopals/set1"
)

func TestChal4(t *testing.T) {
	blockSize := findECBKeySize()

	// Detect ECB
	data := bytes.Repeat([]byte("\x00"), 48)
	cipher := encryptionOracleSame(data)
	if !set1.DetectECB(cipher) {
		t.Fail()
	}

	plain := breakECB(blockSize)

	fmt.Println(string(plain))
}
