package set1

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/sespiros/go-cryptopals/util"
)

func DecryptAESECB(cipher, key []byte) (plain []byte) {
	block, err := aes.NewCipher(key)
	util.Check(err)

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

func Chal7() {
	data, err := ioutil.ReadFile("7.txt")
	util.Check(err)

	key := []byte("YELLOW SUBMARINE")

	cipher, _ := base64.StdEncoding.DecodeString(string(data))

	plain := DecryptAESECB(cipher, key)

	fmt.Println(string(plain))
}
