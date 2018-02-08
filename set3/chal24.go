package set3

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/sespiros/go-cryptopals/util"
)

func mtOracle(plain []byte) ([]byte, uint16) {
	msg := make([]byte, 10+rand.Intn(40))
	rand.Read(msg)
	msg = append(msg, plain...)

	keyb := make([]byte, 2)
	rand.Read(keyb)
	key := uint16(keyb[0])<<8 + uint16(keyb[1])

	return mtEncrypt(msg, key), key
}

func mtEncrypt(plain []byte, seed uint16) []byte {
	var gen mt19337
	gen.init(uint32(seed))
	keystream := make([]byte, len(plain)+3)

	for i := 0; i < len(plain); i += 4 {
		x := gen.extractNumber()
		keystream[i] = byte(x)
		keystream[i+1] = byte(x >> 8)
		keystream[i+2] = byte(x >> 16)
		keystream[i+3] = byte(x >> 24)
	}
	keystream = keystream[:len(plain)]

	return util.Xor(plain, keystream)
}

func mtDecrypt(ct []byte, seed uint16) []byte {
	var gen mt19337
	gen.init(uint32(seed))
	keystream := make([]byte, len(ct)+3)

	for i := 0; i < len(ct); i += 4 {
		x := gen.extractNumber()
		keystream[i] = byte(x)
		keystream[i+1] = byte(x >> 8)
		keystream[i+2] = byte(x >> 16)
		keystream[i+3] = byte(x >> 24)
	}
	keystream = keystream[:len(ct)]

	return util.Xor(ct, keystream)
}

func recoverMtSeed(ct, plain []byte) uint16 {
	for s := 0; s < 0xffff; s++ {
		if bytes.HasSuffix(mtDecrypt(ct, uint16(s)), plain) {
			return uint16(s)
		}
	}
	panic("key not found!")
}

func generateMtToken() []byte {
	plain := []byte(";sespiros;password_reset=true")
	msg := make([]byte, 10+rand.Intn(5))
	rand.Read(msg)
	msg = append(msg, plain...)
	ct := mtEncrypt(msg, uint16(time.Now().Unix()))

	return ct
}

func detectMtToken(token []byte) bool {
	for i := 0; i < 60*60*24; i++ {
		if bytes.HasSuffix(mtDecrypt(token, uint16(i)), []byte("true")) {
			return true
		}
	}

	return false
}
