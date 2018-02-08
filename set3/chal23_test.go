package set3

import (
	"fmt"
	"testing"
)

func TestChal23(t *testing.T) {

	var gen, gen2 mt19337
	gen.init(0)
	gen2.init(0)

	for i := 0; i < 624; i++ {
		gen2.mt[i] = untemper(gen.extractNumber())
	}

	for i := 0; i < 1000; i++ {
		if gen.extractNumber() != gen2.extractNumber() {
			fmt.Printf("Error in clone random generation %d\n", i)
			fmt.Println(gen.extractNumber())
			fmt.Println(gen2.extractNumber())
			t.FailNow()
		}
	}
	fmt.Println("Clone successfull!")
}
