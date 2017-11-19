package set3

import (
	"encoding/binary"

	"github.com/sespiros/go-cryptopals/util"
)

func generateNC(nonce, counter int) []byte {
	nb := make([]byte, 8)
	binary.LittleEndian.PutUint64(nb, uint64(nonce))
	cb := make([]byte, 8)
	binary.LittleEndian.PutUint64(cb, uint64(counter))

	return append(nb, cb...)
}

func EncryptAESCTR(plain, key []byte, nonce int) (cipher []byte) {
	var i int
	numBlocks := len(plain) / BLOCKSIZE

	for i = 0; i < numBlocks; i++ {
		nc := generateNC(nonce, i)
		keyStream := util.EncryptAESECB(nc, key)
		cipher = append(cipher, util.Xor(plain[i*BLOCKSIZE:(i+1)*BLOCKSIZE], keyStream)...)
	}

	// Encrypting the rest of the bytes
	nc := generateNC(nonce, i)
	keyStream := util.EncryptAESECB(nc, key)

	for i := 0; i < len(plain)%BLOCKSIZE; i++ {
		cipher = append(cipher, plain[numBlocks*BLOCKSIZE+i]^keyStream[i])
	}

	return cipher
}

func DecryptAESCTR(cipher, key []byte, nonce int) (plain []byte) {
	var i int
	numBlocks := len(cipher) / BLOCKSIZE

	for i = 0; i < numBlocks; i++ {
		nc := generateNC(nonce, i)
		keyStream := util.EncryptAESECB(nc, key)
		plain = append(plain, util.Xor(cipher[i*BLOCKSIZE:(i+1)*BLOCKSIZE], keyStream)...)
	}

	// Encrypting the rest of the bytes
	nc := generateNC(nonce, i)
	keyStream := util.EncryptAESECB(nc, key)

	for i := 0; i < len(cipher)%BLOCKSIZE; i++ {
		plain = append(plain, cipher[numBlocks*BLOCKSIZE+i]^keyStream[i])
	}

	return plain
}
