package set2

import (
	"math/rand"
	"time"

	"github.com/sespiros/go-cryptopals/util"
)

func EncryptionOracle(plain []byte) (cipher []byte, isECB bool) {
	key := util.GenerateRandomKey16()
	rand.Seed(time.Now().UnixNano())

	plain = append(util.RandomBytes(rand.Intn(6)+5), plain...)
	plain = append(plain, util.RandomBytes(rand.Intn(6)+5)...)
	plain = PKCSpadding(plain, len(key))

	mode := rand.Intn(2)
	switch mode {
	case 0:
		isECB = true
		cipher = util.EncryptAESECB(plain, key)
	case 1:
		iv := util.GenerateRandomKey16()
		cipher = EncryptAESCBC(plain, key, iv)
	}

	return cipher, isECB
}
