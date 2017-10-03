package set2

import (
	"bytes"
	"testing"

	"github.com/sespiros/go-cryptopals/set1"
)

func TestChal11(t *testing.T) {
	data := bytes.Repeat([]byte("\x00"), 48)

	cipher, isECB := EncryptionOracle(data)

	if set1.DetectECB(cipher) {
		if !isECB {
			t.Fail()
		}
	} else {
		if isECB {
			t.Fail()
		}
	}
}
