package set2

import "fmt"

type errBadPadding int

func (e errBadPadding) Error() string {
	return fmt.Sprintf("checkPKCS: Bad padding")
}

// PKCSStrip strips PKCS padding from passed data
func PKCSStrip(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return data, nil
	}

	lastByte := data[len(data)-1:][0]
	num := int(lastByte)

	if len(data) < num {
		return data, nil
	}

	for i := 1; i <= num; i++ {
		if data[len(data)-i] != lastByte {
			return nil, errBadPadding(0)
		}
	}

	return data[:len(data)-num], nil
}
