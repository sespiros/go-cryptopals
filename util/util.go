package util

import (
	"crypto/aes"
	crand "crypto/rand"
	"fmt"
	"log"
	"math/rand"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Xor(a, b []byte) []byte {
	if len(a) != len(b) {
		fmt.Println(a)
		fmt.Println(b)
		log.Fatal("Xor: Given strings not of same size")
	}

	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
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
