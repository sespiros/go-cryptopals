package set3

import (
	"fmt"
	"testing"
)

func TestChal21(t *testing.T) {

	var mt mt19337

	mt.init(0)

	fmt.Println(mt.extract_number())

}
