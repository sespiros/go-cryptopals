package set2

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/sespiros/go-cryptopals/util"
)

func TestChal10(t *testing.T) {
	data, err := ioutil.ReadFile("10.txt")
	util.Check(err)

	key := []byte("YELLOW SUBMARINE")
	iv := bytes.Repeat([]byte("\x00"), 16)

	cipher, _ := base64.StdEncoding.DecodeString(string(data))
	DecryptAESCBC(cipher, key, iv)

}
