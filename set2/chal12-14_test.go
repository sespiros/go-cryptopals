package set2

import (
	"testing"
)

func TestChal12(t *testing.T) {
	simple := customECB(16, 0)

	breakCustomECB(simple)
}

func TestChal14(t *testing.T) {
	custom := customECB(16, 32)

	breakCustomECB(custom)
}
