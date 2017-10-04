package set2

import (
	"bytes"
	"strings"

	"github.com/sespiros/go-cryptopals/util"
)

func strip(input string) (output string) {
	output = strings.Replace(input, ";", "%3b", -1)
	output = strings.Replace(output, "=", "%3d", -1)

	return output
}

type Login struct {
	key []byte
}

func (l Login) EncryptUserData(input string) []byte {
	input = strip(input)
	userData := "comment1=cooking%20MCs;userdata=" + input + ";comment2=%20like%20a%20pound%20of%20bacon"
	plain := PKCSPadding([]byte(userData), 16)

	iv := bytes.Repeat([]byte("\x00"), 16)
	enc := EncryptAESCBC(plain, l.key, iv)

	return enc
}

func (l Login) IsAdmin(enc []byte) bool {
	iv := bytes.Repeat([]byte("\x00"), 16)
	dec := DecryptAESCBC(enc, l.key, iv)

	dec, err := PKCSStrip(dec)
	util.Check(err)

	// Ignoring parsing error so that bitflip can continue
	dataObj, _ := Parse(string(dec), ";")

	ret, ok := dataObj["admin"]
	if ok && ret == "true" {
		return true
	}
	return false
}

func bitflip(login Login) []byte {
	input := "aaaaa\x00admin\x00true"

	enc := login.EncryptUserData(input)

	// 00000000  63 6f 6d 6d 65 6e 74 31  3d 63 6f 6f 6b 69 6e 67  |comment1=cooking|
	// 00000010  25 32 30 4d 43 73 3b 75  73 65 72 64 61 74 61 3d  |%20MCs;userdata=|
	// 00000020  61 61 61 61 61 00 61 64  6d 69 6e 00 74 72 75 65  |aaaaa.admin.true|
	// 00000030  3b 63 6f 6d 6d 65 6e 74  32 3d 25 32 30 6c 69 6b  |;comment2=%20lik|
	// 00000040  65 25 32 30 61 25 32 30  70 6f 75 6e 64 25 32 30  |e%20a%20pound%20|
	// 00000050  6f 66 25 32 30 62 61 63  6f 6e 06 06 06 06 06 06  |of%20bacon......|
	// We just need to brute force bits 26 to match ; and 2c to match = (the zeroes in the third line)

	for a := '\x00'; a < 256; a++ {
		for b := '\x00'; b < 256; b++ {
			enc[16+5] = byte(a)
			enc[16+11] = byte(b)
			if login.IsAdmin(enc) {
				return enc
			}
		}
	}

	return enc
}
