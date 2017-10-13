package set3

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/sespiros/go-cryptopals/util"
)

func TestChal17(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	login := Login{util.RandomBytes(16)}

	enc, iv := login.encryptCBC()

	plain := paddingOracle(login, enc, iv)

	fmt.Println(hex.Dump(plain))
}
