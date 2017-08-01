package set2

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/sespiros/go-cryptopals/set1"
	"github.com/sespiros/go-cryptopals/util"
)

func encryptionOracleSame(plain []byte) (cipher []byte) {
	key, err := hex.DecodeString("67a92b6998eaacf13882b181b25c6f7b")
	util.Check(err)

	rand.Seed(time.Now().UnixNano())

	str := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"

	data, err := base64.StdEncoding.DecodeString(str)
	util.Check(err)

	plain = append(plain, data...)
	plain = PKCSpadding(plain, len(key))

	cipher = util.EncryptAESECB(plain, key)

	return cipher
}

func findECBKeySize() int {
	for i := 1; i < 40; i++ {
		data := bytes.Repeat([]byte("\x41"), i)
		cipher := encryptionOracleSame(data)
		if set1.DetectECB(cipher) {
			return i / 2
		}
	}

	return 0
}

func breakBlock(blockNum, blockSize int, known []byte) []byte {
	plain := make([]byte, 0)

	for i := 1; i <= blockSize; i++ {
		blocks := make(map[string]int)

		input := known[i:]
		found := append(input, plain...)

		ind := blockNum * blockSize

		for i := 0; i < 256; i++ {
			data := append(found, byte(i))
			cipher := encryptionOracleSame(data)
			blocks[string(cipher[:blockSize])] = i
		}
		cipher := encryptionOracleSame(input)
		plain = append(plain, byte(blocks[string(cipher[ind:ind+blockSize])]))
	}

	return plain
}

func breakECB(blockSize int) []byte {
	cipherLen := len(encryptionOracleSame(make([]byte, 0)))
	blocks := cipherLen / blockSize
	ret := make([]byte, 0)

	known := bytes.Repeat([]byte("\x41"), blockSize)
	for i := 0; i < blocks; i++ {
		known = breakBlock(i, blockSize, known)
		ret = append(ret, known...)
	}

	return ret
}
