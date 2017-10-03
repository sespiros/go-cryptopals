package set2

import (
	"bytes"
	"encoding/base64"
	"log"
	"math/rand"
	"time"

	"github.com/sespiros/go-cryptopals/set1"
	"github.com/sespiros/go-cryptopals/util"
)

func customECB(keyLen int, numOfRandomBytesMax int) func([]byte) []byte {
	key := util.RandomBytes(keyLen)
	rand.Seed(time.Now().UnixNano())

	const str = "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	target, err := base64.StdEncoding.DecodeString(str)
	util.Check(err)

	var numOfRandomBytes = 0
	if numOfRandomBytesMax != 0 {
		numOfRandomBytes = rand.Intn(numOfRandomBytesMax)
	}

	random := util.RandomBytes(numOfRandomBytes)

	return func(plain []byte) (cipher []byte) {
		plain = append(random, plain...)
		plain = append(plain, target...)
		plain = PKCSPadding(plain, len(key))

		cipher = util.EncryptAESECB(plain, key)

		return cipher
	}
}

func findECBKeySize(encryptECB func([]byte) []byte) int {
	for i := 0; i < 48; i++ {
		data := bytes.Repeat([]byte("\x41"), i)
		cipher := encryptECB(data)
		if set1.DetectECB(cipher) {
			return 16 // DetectECB only detects keysize 16 for now
		}
	}

	return 0
}

func findRandomOffsetLength(encryptECB func([]byte) []byte, keySize int) int {
	flagBlock := bytes.Repeat([]byte("\x41"), keySize)
	var encFlagBlock []byte
	data := append(flagBlock, flagBlock...)
	data = append(data, flagBlock...)

	// Find encoding of flagBlock
	isSame := 1
	cipher := encryptECB(data)
	for i := 0; i < len(cipher)/keySize-1 && isSame != 0; i++ {
		encFlagBlock = cipher[i*keySize : (i+1)*keySize]
		isSame = bytes.Compare(encFlagBlock, cipher[(i+1)*keySize:(i+2)*keySize])
	}

	// Calculate blockOffset by searching for encFlagBlock
	data = flagBlock
	block := 0
	blockOffset := 0
	for block == 0 {
		cipher := encryptECB(data)
		for i := 1; i < len(cipher)/keySize; i++ {
			if bytes.Compare(encFlagBlock, cipher[i*keySize:(i+1)*keySize]) == 0 {
				block = i
			}
		}

		blockOffset++
		data = append([]byte("\x00"), data...)
	}

	offset := block*keySize - blockOffset + 1

	return offset
}

func breakBlock(encryptECB func([]byte) []byte, blockNum, blockSize int, input []byte, padBlock int) []byte {
	blockOffset := blockNum * blockSize
	padBlockOffset := padBlock * blockSize
	var found []byte

	for i := 1; i <= blockSize; i++ {
		blocks := make(map[string]byte)

		input = input[1:]

		// Build dictionary
		for j := 0; j < 256; j++ {
			data := append(input, found...)
			data = append(data, byte(j))
			cipher := encryptECB(data)
			blocks[string(cipher[padBlockOffset:padBlockOffset+blockSize])] = byte(j)
		}

		// Find match
		cipher := encryptECB(input)
		block := string(cipher[blockOffset : blockOffset+blockSize])

		nextByte, ok := blocks[block]

		// If reached end of plaintext, fix PKCS padding
		if !ok && found != nil {
			found = PKCSPadding(found[:len(found)-1], blockSize)
			break
		}
		found = append(found, nextByte)
	}

	return found
}

func breakCustomECB(encryptECB func([]byte) []byte) {
	blockSize := findECBKeySize(encryptECB)

	// Detect ECB
	data := bytes.Repeat([]byte("\x41"), 3*blockSize)
	cipher := encryptECB(data)
	if !set1.DetectECB(cipher) {
		log.Fatal("breakCustomECB: Didn't detect ECB function")
	}

	// Find offsetLength
	offsetLen := findRandomOffsetLength(encryptECB, blockSize)
	offsetPadLen := blockSize - offsetLen%blockSize
	if offsetLen%blockSize == 0 {
		offsetPadLen = 0
	}

	pad := bytes.Repeat([]byte("\x00"), offsetPadLen)

	startBlock := (offsetLen + blockSize - 1) / blockSize

	cipherLen := len(encryptECB(pad))
	blockNum := cipherLen / blockSize

	// Dont try to break a block that is only padding
	t := append(pad, '\x00')
	if len(encryptECB(t)) > cipherLen {
		blockNum--
	}

	var plain []byte

	known := append(pad, bytes.Repeat([]byte("\x41"), blockSize)...)
	for i := startBlock; i < blockNum; i++ {
		known = breakBlock(encryptECB, i, blockSize, known, startBlock)
		plain = append(plain, known...)
		known = append(pad, known...)
	}

	plain, err := PKCSStrip(plain)
	util.Check(err)
}
