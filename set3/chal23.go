package set3

func untemper(mt []int) []int {

	for i := 623; i >= 0; i-- {
		// Inverse of y = y ^ (y >> 18) explanation 18+18 > 32
		mt[i] = mt[i] ^ mt[i]>>18

		// Inverse of y = y ^ ((y << 15) & 0xefc60000)
		mt[i] = mt[i] ^ ((mt[i] & (0xefc60000 >> 15)) << 15)

		// Inverse of y = y ^ ((y << 7) & 0x9d2c5680)
		t := mt[i]
		t = ((t & 0x0000002d) << 7) ^ mt[i]
		t = ((t & 0x000018ad) << 7) ^ mt[i]
		t = ((t & 0x001a58ad) << 7) ^ mt[i]
		mt[i] = ((t & 0x013a58ad) << 7) ^ mt[i]

		// Inverse of y = y ^ (y >> 11)
		// need to fully implement the inverse xor because 11+11 < 32
		top := mt[i] & 0xffe00000
		mid := mt[i] & 0x001ffc00
		low := mt[i] & 0x000003ff

		mt[i] = top | ((top >> 11) ^ mid) | ((((top >> 11) ^ mid) >> 11) ^ low)
	}

	return mt
}
