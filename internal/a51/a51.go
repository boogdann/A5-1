package a51

import (
	"2/internal/bytes"
	"fmt"
)

const (
	SizeReg1 = 19
	SizeReg2 = 22
	SizeReg3 = 23

	Method1 = iota
	Method2

	bit7  = 2 << 7
	bit8  = 2 << 8
	bit10 = 2 << 10
	bit13 = 2 << 13
	bit16 = 2 << 16
	bit17 = 2 << 17
	bit18 = 2 << 18
	bit19 = 2 << 19
	bit20 = 2 << 20
	bit21 = 2 << 21
	bit22 = 2 << 22
)

type A51 struct {
	r1 uint64
	r2 uint64
	r3 uint64
}

func New() *A51 {
	return &A51{}
}

func (a *A51) InitRegs(method int, key uint64) error {
	a.r1 = 0
	a.r2 = 0
	a.r3 = 0

	switch method {
	case Method1:
		a.initRegsMethod1(key)
	case Method2:
		a.initRegsMethod2(key)
	default:
		return fmt.Errorf("invalid method")
	}

	return nil
}

func (a *A51) GenerateKeyStream(length int) []byte {
	fmt.Printf("________________________________________________________________________\n")

	var bit uint8
	keyStream := make([]byte, length)
	for i := 0; i <= length; i++ {
		bit = uint8(((a.r1&(2<<SizeReg1))>>SizeReg1 + 1) ^
			((a.r2&(2<<SizeReg2))>>SizeReg2 + 1) ^
			((a.r3&(2<<SizeReg3))>>SizeReg3 + 1),
		)

		keyStream[i] = bit
		a.shiftRegisters()
		length--
	}
	return keyStream
}

func (a *A51) initRegsMethod1(key uint64) {
	var bit uint8
	for i := 0; i < 64; i++ {
		bit = uint8((key >> i) & 1)

		//fmt.Printf("%064b\n", a.r1)
		a.r1 = a.xorLastBit(a.r1, bit)
		a.r2 = a.xorLastBit(a.r2, bit)
		a.r3 = a.xorLastBit(a.r3, bit)

		a.shiftRotateLeft()
		fmt.Printf("%3d: %064b\n", i, a.r1)
		//fmt.Printf("%3d: %064b\n", i, a.r2)
		//fmt.Printf("%3d: %064b\n", i, a.r3)
	}

	fmt.Printf("ar1: %064b\n", a.r1)
	fmt.Printf("ar2: %064b\n", a.r2)
	fmt.Printf("ar3: %064b\n", a.r3)

	for i := 0; i < 100; i++ {
		a.shiftRegisters()
	}

	fmt.Printf("ar1: %064b\n", a.r1)
	fmt.Printf("ar2: %064b\n", a.r2)
	fmt.Printf("ar3: %064b\n", a.r3)
}

func (a *A51) shiftRegistersWithoutBit8() {
	x := uint8((a.r1 & bit8) >> 9)
	y := uint8((a.r2 & bit10) >> 11)
	z := uint8((a.r3 & bit10) >> 11)

	f := uint64((x & y) | (x & z) | (y & z))

	var bit uint8
	t := (a.r1 & bit8) >> 9
	_ = t
	if f == ((a.r1 & bit8) >> 9) {
		bit = uint8(((a.r1 & bit13) >> 14) ^ ((a.r1 & bit16) >> 17) ^
			((a.r1 & bit17) >> 18) ^ ((a.r1 & bit18) >> 19))
		a.r1 = bytes.RotateLeftWithoutBit(a.r1, 1, 8, SizeReg1)
		a.r1 = a.xorLastBit(a.r1, bit)
	}

	if f == ((a.r2 & bit10) >> 11) {
		bit = uint8(((a.r2 & bit20) >> 21) ^
			((a.r2 & bit21) >> 22))
		a.r2 = bytes.RotateLeftWithoutBit(a.r2, 1, 10, SizeReg2)
		a.r2 = a.xorLastBit(a.r2, bit)
	}

	if f == ((a.r3 & bit10) >> 11) {
		bit = uint8(((a.r3 & bit7) >> 8) ^ ((a.r3 & bit20) >> 21) ^
			((a.r3 & bit21) >> 22) ^ ((a.r3 & bit22) >> 23))
		a.r3 = bytes.RotateLeftWithoutBit(a.r3, 1, 10, SizeReg3)
		a.r3 = a.xorLastBit(a.r3, bit)
	}
}

func (a *A51) shiftRegisters() {
	var bit uint8
	bit = uint8(((a.r1 & bit13) >> 14) ^ ((a.r1 & bit16) >> 17) ^
		((a.r1 & bit17) >> 18) ^ ((a.r1 & bit18) >> 19))
	a.r1 = bytes.RotateLeftWithoutBit(a.r1, 1, 8, SizeReg1)
	a.r1 = a.xorLastBit(a.r1, bit)

	bit = uint8(((a.r2 & bit20) >> 21) ^
		((a.r2 & bit21) >> 22))
	a.r2 = bytes.RotateLeftWithoutBit(a.r2, 1, 10, SizeReg2)
	a.r2 = a.xorLastBit(a.r2, bit)

	bit = uint8(((a.r3 & bit7) >> 8) ^ ((a.r3 & bit20) >> 21) ^
		((a.r3 & bit21) >> 22) ^ ((a.r3 & bit22) >> 23))
	a.r3 = bytes.RotateLeftWithoutBit(a.r3, 1, 10, SizeReg3)
	a.r3 = a.xorLastBit(a.r3, bit)
}

func (a *A51) initRegsMethod2(key uint64) {
	a.r1 = 0
	a.r2 = 0
	a.r3 = 0

	for i := 0; i < 64; i++ {
		bit := (key >> uint(i)) & 1
		a.r1 = a.r1 ^ (bit ^ ((a.r1 >> 1) & 1) ^ ((a.r1 >> 2) & 1))
		//} else {
		//	if i < 20 {
		//		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2-10), data[i])
		//		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2-10), key[i])
		//		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2-10), cryptData[i])
		//	} else {
		//		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+2-20), data[i])
		//		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+2-20), key[i])
		//		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+2-20), cryptData[i])
		//	}
		//}	a.r2 = a.r2 ^ (bit ^ ((a.r2 >> 1) & 1) ^ ((a.r2 >> 2) & 1))
		a.r3 = a.r3 ^ (bit ^ ((a.r3 >> 1) & 1) ^ ((a.r3 >> 2) & 1))
		a.shiftRegisters()
	}

	for i := 0; i < 222; i++ {
		a.shiftRotateLeft()
	}
}

func (a *A51) shiftRotateLeft() {
	a.r1 = bytes.RotateLeft(a.r1, 1, SizeReg1)
	a.r2 = bytes.RotateLeft(a.r2, 1, SizeReg2)
	a.r3 = bytes.RotateLeft(a.r3, 1, SizeReg3)
}

func (a *A51) xorLastBit(num uint64, bit uint8) uint64 {
	mask := uint64(1) << 0
	//fmt.Printf("%b\n", mask)
	res := num | (uint64(bit) & mask)
	//fmt.Printf("%b\n", res)

	return res
}
