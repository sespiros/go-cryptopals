package set3

func _int32(x int) int {
	// Get the 32 least significant bits.
	return int(0xFFFFFF & x)
}

type mt19337 struct {
	index int
	mt    []int
}

func (self *mt19337) init(seed int) {
	// Initialize the index to 0
	self.index = 0
	self.mt = make([]int, 624)

	self.mt[0] = seed // Initialize the initial state to the seed

	for i := 1; i < 624; i++ {
		self.mt[i] = _int32(1812433253*(self.mt[i-1]^self.mt[i-1]>>30) + i)
	}
}

func (self *mt19337) extract_number() int {
	if self.index >= 624 {
		self.twist()
	}

	y := self.mt[self.index]

	// Right shift by 11 bits
	y = y ^ y>>11
	// Shift y left by 7 and take the bitwise and of 2636928640
	y = y ^ y<<7&0x9d2c5680
	// Shift y left by 15 and take the bitwise and of y and 4022730752
	y = y ^ y<<15&0xefc60000
	// Right shift by 18 bits
	y = y ^ y>>18

	self.index = self.index + 1

	return _int32(y)
}

func (self *mt19337) twist() {
	for i := 0; i < 624; i++ {
		// Get the most significant bit and add it to the less significant
		// bits of the next number
		y := _int32((self.mt[i] & 0x80000000) + (self.mt[(i+1)%624] & 0x7fffffff))
		self.mt[i] = self.mt[(i+397)%624] ^ y>>1

		if y%2 != 0 {
			self.mt[i] = self.mt[i] ^ 0x9908b0df
		}
	}

	self.index = 0
}
