package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Chal1() {
	const t = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hex, _ := hex.DecodeString(t)
	encoded := base64.StdEncoding.EncodeToString(hex)
	fmt.Println(encoded)
}
