package set3

func untemper(y uint32) uint32 {

	// Inverse of y = y ^ (y >> 18) explanation 18+18 > 32
	y ^= y >> 18

	// // Inverse of y = y ^ ((y << 15) & 0xefc60000)
	y ^= y << 15 & 4022730752

	// Inverse of y = y ^ ((y << 7) & 0x9d2c5680)
	for i := 0; i < 7; i++ {
		y ^= y << 7 & 2636928640
	}

	// Inverse of y = y ^ (y >> 11)
	y ^= y>>11 ^ y>>(11*2)

	return y
}
