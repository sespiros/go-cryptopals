package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/sespiros/go-cryptopals/util"
)

func hamming(s1, s2 []byte) int {
	if len(s1) != len(s2) {
		log.Fatal("hamming: strings have different size")
	}

	count64 := func(x uint64) uint64 {
		x = (x & 0x5555555555555555) + ((x & 0xAAAAAAAAAAAAAAAA) >> 1)
		x = (x & 0x3333333333333333) + ((x & 0xCCCCCCCCCCCCCCCC) >> 2)
		x = (x & 0x0F0F0F0F0F0F0F0F) + ((x & 0xF0F0F0F0F0F0F0F0) >> 4)
		x *= 0x0101010101010101
		return ((x >> 56) & 0xFF)
	}

	distance := 0

	for i := range s1 {
		r := s1[i] ^ s2[i]
		distance += int(count64(uint64(r)))
	}

	return distance
}

type result struct {
	ks   int
	dist float32
}

type ByDist []result

func (a ByDist) Len() int           { return len(a) }
func (a ByDist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDist) Less(i, j int) bool { return a[i].dist < a[j].dist }

func _findKeyFitness(enc []byte, ks int, ch chan result) {
	avgDist := float32(0)
	iter := 0

	for i := 0; i+2*ks < len(enc); i += ks {
		a, b := enc[i:i+ks], enc[i+ks:i+2*ks]
		dist := hamming(a, b)
		avgDist += float32(dist) / float32(ks)
		iter++
	}

	dist := avgDist / float32(iter)

	ch <- result{ks, dist}
}

func guessKeySize(enc []byte) int {
	c := make(chan result)
	const start = 16
	const end = 50
	results := make([]result, 0, end-start)

	for ks := start; ks <= end; ks++ {
		go _findKeyFitness(enc, ks, c)
	}

	for ks := start; ks <= end; ks++ {
		results = append(results, <-c)
	}

	sort.Sort(ByDist(results))

	return results[0].ks
}

func split(enc []byte, ks int) [][]byte {
	blockCount := len(enc) / ks
	blocks := make([][]byte, blockCount)
	for i := 0; i < blockCount; i++ {
		blocks[i] = enc[i*ks : (i+1)*ks]
	}
	return blocks
}

func transpose(blocks [][]byte) [][]byte {
	tblocks := make([][]byte, len(blocks[0]))
	for i := range tblocks {
		tblocks[i] = make([]byte, len(blocks))
		for j := range blocks {
			tblocks[i][j] = blocks[j][i]
		}
	}
	return tblocks
}

func BreakRepeatedKeyXor(enc []byte) []byte {
	ks := guessKeySize(enc)
	blocks := split(enc, ks)
	blocks = transpose(blocks)
	var key []byte
	for _, block := range blocks {
		_, _, keyChar := BreakSingleCharXor(hex.EncodeToString(block))
		key = append(key, byte(keyChar))
	}

	// fmt.Println("The key is")
	// fmt.Println(string(key))

	plain := RepeatedXor(enc, key)

	return plain
}

func Chal6() {
	data, err := ioutil.ReadFile("6.txt")
	util.Check(err)

	enc, err := base64.StdEncoding.DecodeString(string(data))

	plain := BreakRepeatedKeyXor(enc)

	fmt.Println("and the plaintext is")
	fmt.Print(string(plain))
}
