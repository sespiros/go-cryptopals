package set3

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sespiros/go-cryptopals/util"
)

func checkLines(clines [][]byte, key []byte) (bool, int) {
	var ok bool

	for i, l := range clines {
		pl := util.Xor(key, l)
		fmt.Printf("%d. ", i)
		for i := 0; i < len(l); i++ {
			if byte(key[i]) != 0 {
				fmt.Printf("%c", pl[i])
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Does this look ok?(y/n) ")
	text, _ := reader.ReadString('\n')

	if text == "y\n" || text == "\n" {
		ok = true
	}

	var num int

	fmt.Printf("Which string next? ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		num = 0
	}

	return ok, num
}

func breakCTR(clines [][]byte) []byte {
	var max int
	for _, l := range clines {
		if len(l) > max {
			max = len(l)
		}
	}

	str := 0
	i := 0
	key := bytes.Repeat([]byte("\x00"), max)

	for {
		fmt.Printf("Guess letter for string %d position %d: ", str, i)
		reader := bufio.NewReader(os.Stdin)
		text, ret := reader.ReadString('\n')
		if ret == io.EOF {
			fmt.Println("Invalid input")
			break
		}
		key[i] = []byte(text)[0] ^ clines[str][i]

		var answer bool
		if answer, str = checkLines(clines, key); answer {
			i++
		}

		if i == max {
			break
		}
	}

	return key
}

func Run() {
	data, err := ioutil.ReadFile("set3/19.txt")
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

	keyr := breakCTR(clines)

	fmt.Println(keyr)

}
