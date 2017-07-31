package set1

import (
	"encoding/hex"
	"fmt"
)

func Xor(a, b []byte) []byte {
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}

func Chal2() {
	const a = "1c0111001f010100061a024b53535009181c"
	const b = "686974207468652062756c6c277320657965"
	ah, _ := hex.DecodeString(a)
	bh, _ := hex.DecodeString(b)

	res := Xor(ah, bh)

	fmt.Println(hex.EncodeToString(res))
}
