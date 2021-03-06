package set2

import (
	"bytes"
	"fmt"
	"testing"
)

func TestChal9(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")
	want := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	output := PKCSPadding(input, 20)
	if !bytes.Equal(output, want) {
		fmt.Println(output)
		fmt.Println(want)
		t.Fail()
	}

}
