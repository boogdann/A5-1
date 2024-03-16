package bytes

//func RotateLeftWithoutBit(num uint64, shift int, bitIndex int, size int) uint64 {
//	num = 143
//	fmt.Printf("Num         %019b\n", num)
//	fmt.Printf("____________\n")
//	shift = 1
//	bitIndex = 7
//	size = 19
//
//	shift = shift % size
//	bit := (num & (2 << (bitIndex - 1))) >> bitIndex
//	rightBits := num & (1 << (uint(bitIndex) - 1))
//	fmt.Printf("Right Bits: %019b\n", rightBits)
//	num = (num << shift) | (rightBits >> (size - shift))
//	fmt.Printf("Num         %019b\n", num)
//	bit = (bit << shift) | (bit >> (1 - shift))
//	fmt.Printf("Bit         %019b\n", bit)
//	num = (num &^ (1 << uint(bitIndex))) | (bit << uint(bitIndex))
//	fmt.Printf("Num         %019b\n", num)
//	return num
//}

func RotateLeftWithoutBit(num uint64, shift int, bitIndex int, size int) uint64 {
	num = ((num << shift) | (num >> (size - shift))) & (uint64(1<<size - 1))
	return SwapAdjacentBits(num, bitIndex-1) & (uint64(1<<size - 1))
}

func SwapAdjacentBits(num uint64, i int) uint64 {
	bitI := (num >> i) & 1
	bitIPlus1 := (num >> (i + 1)) & 1

	maskI := bitI << (i + 1)
	maskIPlus1 := bitIPlus1 << i

	num &^= (1 << i) | (1 << (i + 1))

	num |= maskI | maskIPlus1

	return num
}

func RotateLeft(num uint64, shift int, size int) uint64 {
	shift = shift % size
	return ((num << shift) | (num >> (size - shift))) & (uint64(1<<size - 1))
}
