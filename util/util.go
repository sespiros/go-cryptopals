package util

import (
	"bytes"
	"crypto/aes"
	crand "crypto/rand"
	"log"
	"math/rand"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Xor(a, b []byte) []byte {
	max := a
	min := b
	if len(b) > len(a) {
		max = b
		min = a
	}

	res := make([]byte, len(max))
	diff := len(max) - len(min)
	min = append(min, bytes.Repeat([]byte("\x00"), diff)...)
	for i := 0; i < len(max); i++ {
		res[i] = min[i] ^ max[i]
	}
	return res
}

func EncryptAESECB(plain, key []byte) (cipher []byte) {
	block, err := aes.NewCipher(key)
	Check(err)

	bs := block.BlockSize()

	cipher = make([]byte, len(plain))
	cblock := cipher
	for len(plain) > 0 {
		block.Encrypt(cblock, plain[:bs])
		cblock = cblock[bs:]
		plain = plain[bs:]
	}

	return cipher
}

func DecryptAESECB(cipher, key []byte) (plain []byte) {
	block, err := aes.NewCipher(key)
	Check(err)

	bs := block.BlockSize()

	plain = make([]byte, len(cipher))
	pblock := plain
	for len(cipher) > 0 {
		block.Decrypt(pblock, cipher[:bs])
		pblock = pblock[bs:]
		cipher = cipher[bs:]
	}

	return plain
}

func GenerateRandomKey16() []byte {
	key := make([]byte, 16)
	_, err := crand.Read(key)
	Check(err)

	return key
}

func RandomBytes(n int) []byte {
	pad := make([]byte, n)
	_, err := rand.Read(pad)
	Check(err)

	return pad
}
