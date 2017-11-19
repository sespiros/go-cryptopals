package set3

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/sespiros/go-cryptopals/util"
)

func TestChal18(t *testing.T) {
	str := "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="
	key := []byte("YELLOW SUBMARINE")
	nonce := 0

	cipher, err := base64.StdEncoding.DecodeString(str)
	util.Check(err)

	// cipher = EncryptAESCTR(key, key, nonce)
	plain := DecryptAESCTR(cipher, key, nonce)

	fmt.Println(hex.Dump(plain))
}
