package set2

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/sespiros/go-cryptopals/set1"
)

func TestChal3(t *testing.T) {
	data := bytes.Repeat([]byte("\x00"), 48)

	cipher, isECB := EncryptionOracle(data)

	if set1.DetectECB(cipher) {
		if isECB {
			fmt.Println("Correctly identified ECB")
		} else {
			t.Fail()
		}
	} else {
		if !isECB {
			fmt.Println("Correctly identified CBC")
		} else {
			t.Fail()
		}
	}
}
