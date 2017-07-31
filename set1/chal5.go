package set1

import (
	"encoding/hex"
	"fmt"
)

func RepeatedXor(s []byte, key []byte) []byte {
	res := make([]byte, len(s))
	for i := range s {
		res[i] = s[i] ^ key[i%len(key)]
	}

	return res
}

func Chal5() {
	const plain = "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"

	fmt.Println(hex.EncodeToString(RepeatedXor([]byte(plain), []byte("ICE"))))

}
