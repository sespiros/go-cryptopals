package set2

func PCKSpadding(data []byte, size int) []byte {
	padSize := size - len(data)
	for i := 0; i < padSize; i++ {
		data = append(data, 4)
	}

	return data
}
