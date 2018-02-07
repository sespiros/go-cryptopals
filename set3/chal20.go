package set3

import (
	"bytes"

	"github.com/sespiros/go-cryptopals/set1"
)

func breakCTRstat(clines [][]byte) []byte {
	for i, _ := range clines {
		clines[i] = clines[i][:48]
	}

	enc := bytes.Join(clines, []byte(""))

	plain := set1.BreakRepeatedKeyXor(enc)

	return plain
}
