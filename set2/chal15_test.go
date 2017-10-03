package set2

import (
	"bytes"
	"testing"
)

func TestChal15(t *testing.T) {
	correct := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	wrong1 := []byte("YELLOW SUBMARINE\x05\x05\x05\x05")
	wrong2 := []byte("YELLOW SUBMARINE\x01\x02\x03\x04")
	stripped := []byte("YELLOW SUBMARINE")

	output, err := PKCSStrip(correct)
	if err != nil {
		if bytes.Compare(output, stripped) != 0 {
			t.Errorf("checkPKCS: mismatch %s %s", string(output), string(stripped))
		}
	}

	output, err = PKCSStrip(wrong2)
	if err == nil {
		t.Errorf("checkPKCS: %s should produce error", string(wrong1))
	}

	output, err = PKCSStrip(wrong2)
	if err == nil {
		t.Errorf("checkPKCS: %s should produce error", string(wrong2))
	}

}
