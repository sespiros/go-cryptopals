package set2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/sespiros/go-cryptopals/util"
)

func TestChal2(t *testing.T) {
	data, err := ioutil.ReadFile("10.txt")
	util.Check(err)

	key := []byte("YELLOW SUBMARINE")
	IV := bytes.Repeat([]byte("\x00"), 16)

	plain := DecryptAESCBC(data, key, IV)

	fmt.Println(string(plain))
}
