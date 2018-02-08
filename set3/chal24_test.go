package set3

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
)

func TestChal24(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAA")
	ct, key := mtOracle(msg)

	if key == recoverMtSeed(ct, msg) {
		fmt.Println("Successfull!")
	}

	token := generateMtToken()
	token2 := make([]byte, 40+rand.Intn(5))
	rand.Read(token2)

	fmt.Println(hex.Dump(token))
	fmt.Println(hex.Dump(token2))

	if !detectMtToken(token) {
		t.Error("Didn't detect token generated with MT13997")
	}

	if detectMtToken(token2) {
		t.Error("Detected token not generated with MT13997")
	}

}
