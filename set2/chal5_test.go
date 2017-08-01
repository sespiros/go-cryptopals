package set2

import (
	"fmt"
	"testing"

	"github.com/sespiros/go-cryptopals/util"
)

func TestChal5(t *testing.T) {
	key := util.GenerateRandomKey16()

	email := "foo@bar.AABBB"
	enc := encryptProfile(email, key)

	// Padding accounting for the strings email= etc
	email = string(PKCSpadding([]byte("foo@bar.AAadmin"), 26))
	cut := encryptProfile(email, key)

	copy(enc[len(key)*2:], cut[len(key):len(key)*2])

	fmt.Println(decryptParseProfile(enc, key))
}
