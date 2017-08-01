package set1

import (
	"encoding/hex"
	"fmt"

	"github.com/sespiros/go-cryptopals/util"
)

func Chal2() {
	const a = "1c0111001f010100061a024b53535009181c"
	const b = "686974207468652062756c6c277320657965"
	ah, _ := hex.DecodeString(a)
	bh, _ := hex.DecodeString(b)

	res := util.Xor(ah, bh)

	fmt.Println(hex.EncodeToString(res))
}
