package set3

import (
	"fmt"
	"testing"
)

func TestChal23(t *testing.T) {

	var gen mt19337
	gen.init(0)

	outputs := make([]int, 624)

	for i := 0; i < 624; i++ {
		outputs[i] = gen.extract_number()
	}

	var gen2 mt19337
	gen2.mt = untemper(outputs)
	gen2.index = 624

	fmt.Println(gen.extract_number())
	fmt.Println(gen2.extract_number())
}
