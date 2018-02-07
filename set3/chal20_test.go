package set3

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/sespiros/go-cryptopals/util"
)

func TestChal20(t *testing.T) {
	data, err := ioutil.ReadFile("20.txt")
	util.Check(err)
	b64lines := strings.Split(string(data), "\n")
	clines := make([][]byte, len(b64lines))

	key := []byte("YELLOW SUBMARINE")
	nonce := 0

	for i, e := range b64lines {
		plain, err := base64.StdEncoding.DecodeString(e)
		util.Check(err)
		cipher := EncryptAESCTR(plain, key, nonce)
		clines[i] = cipher
	}

	plain := breakCTRstat(clines)

	fmt.Println(string(plain))

}
