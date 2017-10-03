package set2

import (
	"testing"

	"github.com/sespiros/go-cryptopals/util"
)

func TestChal13(t *testing.T) {
	key := util.GenerateRandomKey16()
	keyLen := len(key)

	email := "foo@bar.AABBB"
	enc := encryptProfile(email, key)

	// 0-------------16 17------------31 32------------48
	// email=foo@bar.AA BBB&uid=10&role= user000000000000 <----------------|
	// email=foo@bar.AA admin00000000000 &uid=10&role=use r000000000000000 |
	//                  extra padding ^ to match this ---------------------|
	email = string(PKCSPadding([]byte("foo@bar.AAadmin"), keyLen+10))
	cut := encryptProfile(email, key)

	copy(enc[keyLen*2:], cut[keyLen:keyLen*2])

	result := decryptParseProfile(enc, key)
	if result["role"] != "admin" {
		t.Fail()
	}
}
