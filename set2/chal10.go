package set2

import (
	"github.com/sespiros/go-cryptopals/util"
)

// EncryptAESCBC Implements aes-cbc encryption
func EncryptAESCBC(plain, key, iv []byte) (cipher []byte) {
	bs := len(key)

	cipher = make([]byte, len(plain))
	cblock := cipher
	prevCblock := iv
	for len(plain) > 0 {
		enc := util.EncryptAESECB(util.Xor(plain[:bs], prevCblock), key)
		copy(cblock[:bs], enc)
		prevCblock = cblock[:bs]
		cblock = cblock[bs:]
		plain = plain[bs:]
	}

	return cipher
}

// DecryptAESCBC Implements aes-cbc decryption
func DecryptAESCBC(cipher, key, iv []byte) (plain []byte) {
	bs := len(key)

	plain = make([]byte, len(cipher))
	pblock := plain
	prevCblock := iv
	for len(cipher) > 0 {
		dec := util.Xor(util.DecryptAESECB(cipher[:bs], key), prevCblock)
		copy(pblock[:bs], dec)
		prevCblock = cipher[:bs]
		pblock = pblock[bs:]
		cipher = cipher[bs:]
	}

	return plain
}
