package set1

import (
	"bufio"
	"os"

	"github.com/sespiros/go-cryptopals/util"
)

func DetectECB(s []byte) bool {
	keylengths := []int{16}

	for _, ks := range keylengths {
		blockCount := len(s) / ks
		blocks := make([][]byte, blockCount)
		hist := make(map[string]int)
		for i := 0; i < blockCount; i++ {
			blocks[i] = s[i*ks : (i+1)*ks]
			hist[string(blocks[i])]++
		}

		for _, v := range hist {
			if v > 1 {
				return true
			}
		}
	}

	return false
}

func Chal8() {
	file, err := os.Open("8.txt")
	util.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go DetectECB([]byte(scanner.Text()))
	}

}
