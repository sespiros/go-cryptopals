package set3

import (
	"bytes"
	"encoding/base64"
	"math/rand"

	"github.com/sespiros/go-cryptopals/set2"
	"github.com/sespiros/go-cryptopals/util"
)

const BLOCKSIZE = 16

type Login struct {
	key []byte
}

func (l Login) encryptCBC() (enc []byte, iv []byte) {
	iv = util.RandomBytes(BLOCKSIZE)

	var strings []string
	strings = append(strings, "MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=")
	strings = append(strings, "MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=")
	strings = append(strings, "MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==")
	strings = append(strings, "MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==")
	strings = append(strings, "MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl")
	strings = append(strings, "MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==")
	strings = append(strings, "MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==")
	strings = append(strings, "MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=")
	strings = append(strings, "MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=")
	strings = append(strings, "MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93")

	str := strings[rand.Intn(len(strings)-1)]
	plain, err := base64.StdEncoding.DecodeString(str)
	util.Check(err)

	enc = set2.EncryptAESCBC(set2.PKCSPadding(plain, BLOCKSIZE), l.key, iv)

	return enc, iv
}

func (l Login) hasValidPadding(enc, iv []byte) bool {
	dec := set2.DecryptAESCBC(enc, l.key, iv)

	dec, err := set2.PKCSStrip(dec)
	if err != nil {
		return false
	}

	return true
}

func paddingOracle(login Login, enc, iv []byte) (plain []byte) {
	numBlocks := len(enc) / len(iv)
	for i := numBlocks; i > 0; i-- {
		plain = append(breakBlock(enc[:i*BLOCKSIZE], iv, login), plain...)
	}

	return plain
}

func breakBlock(enc, iv []byte, login Login) (plain []byte) {
	pad := bytes.Repeat([]byte{'\x00'}, BLOCKSIZE)

	for i := 0; i < BLOCKSIZE; i++ {
		knownBytes := plain
		padBytes := bytes.Repeat([]byte{byte(i + 1)}, i+1)

		for j := 0; j < 256; j++ {
			tampered := append(iv, enc...)

			known := append(knownBytes, pad...)
			rest := bytes.Repeat([]byte{byte(j)}, len(enc)-len(knownBytes))
			known = append(rest, known...)

			valid := append(padBytes, pad...)

			tampered = util.Xor(tampered, known)
			tampered = util.Xor(tampered, valid)

			if login.hasValidPadding(tampered[len(iv):], tampered[:len(iv)]) {
				plain = append([]byte{byte(j)}, plain...)
				break
			}
		}
	}

	return plain

}
