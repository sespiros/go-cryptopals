package set1

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/sespiros/go-cryptopals/util"
)

func Chal7() {
	data, err := ioutil.ReadFile("7.txt")
	util.Check(err)

	key := []byte("YELLOW SUBMARINE")

	cipher, _ := base64.StdEncoding.DecodeString(string(data))

	plain := util.DecryptAESECB(cipher, key)

	fmt.Println(string(plain))
}
