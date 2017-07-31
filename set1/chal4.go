package set1

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sespiros/go-cryptopals/util"
)

func Chal4() {
	var best string
	max := 0.0

	file, err := os.Open("4.txt")
	util.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		score, str, _ := BreakSingleCharXor(scanner.Text())
		if score > max {
			max = score
			best = str
		}
	}

	util.Check(scanner.Err())

	fmt.Printf("\"%v\" with score %v", strings.TrimSpace(best), max)
}
