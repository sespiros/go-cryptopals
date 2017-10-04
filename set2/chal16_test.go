package set2

import (
	"math/rand"
	"testing"
	"time"

	"github.com/sespiros/go-cryptopals/util"
)

func TestChal16(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	login := Login{util.RandomBytes(16)}

	enc := bitflip(login)

	if !login.IsAdmin(enc) {
		t.Fail()
	}
}
