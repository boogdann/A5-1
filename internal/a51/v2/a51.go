package a51v2

import "fmt"

const (
	Method1 = 1
	Method2 = 2

	SizeReg1 = 19
	SizeReg2 = 22
	SizeReg3 = 23

	SizeKey = 64

	PosSync1 = 8
	PosSync2 = 10
	PosSync3 = 10
)

type A51 struct {
	r1 []byte
	r2 []byte
	r3 []byte
}

func New() *A51 {
	return &A51{}
}

func (a *A51) InitRegs(method int, key []byte) error {
	a.r1 = make([]byte, SizeReg1)
	a.r2 = make([]byte, SizeReg2)
	a.r3 = make([]byte, SizeReg3)

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
	var bit uint8
	keyStream := make([]byte, length+1)
	for i := 0; i <= length; i++ {
		a.shiftRegsWithSyncBit()
		bit = a.r1[0] ^ a.r2[0] ^ a.r3[0]

		keyStream[i] = bit
	}
	return keyStream
}

func (a *A51) initRegsMethod1(key []byte) {
	var idx int
	for i := 0; i < SizeKey; i++ {

		idx = SizeKey - i - 1
		a.r1[SizeReg1-1] = a.r1[SizeReg1-1] ^ key[idx]
		a.r2[SizeReg2-1] = a.r2[SizeReg2-1] ^ key[idx]
		a.r3[SizeReg3-1] = a.r3[SizeReg3-1] ^ key[idx]

		a.shiftRegsWithoutSyncBit(key[idx], key[idx], key[idx])
	}

	for i := 0; i < 100; i++ {
		a.shiftRegsWithSyncBit()
	}
	var i = 0
	_ = i
}

func (a *A51) shiftRegsWithoutSyncBit(bit1, bit2, bit3 byte) {
	var setBit byte

	setBit = a.r1[SizeReg1-13-1] ^ a.r1[SizeReg1-16-1] ^ a.r1[SizeReg1-17-1] ^ a.r1[SizeReg1-18-1]
	a.r1 = rotateLeft(a.r1)
	a.r1[len(a.r1)-1] ^= setBit

	setBit = a.r2[SizeReg2-20-1] ^ a.r2[SizeReg2-21-1]
	a.r2 = rotateLeft(a.r2)
	a.r2[len(a.r2)-1] ^= setBit

	setBit = a.r3[SizeReg3-7-1] ^ a.r3[SizeReg3-20-1] ^ a.r3[SizeReg3-21-1] ^ a.r3[SizeReg3-22-1]
	a.r3 = rotateLeft(a.r3)
	a.r3[len(a.r3)-1] ^= setBit
}

func (a *A51) shiftRegsWithSyncBit() {
	x := a.r1[SizeReg1-PosSync1-1]
	y := a.r2[SizeReg2-PosSync2-1]
	z := a.r3[SizeReg3-PosSync3-1]

	f := (x & y) | (x & z) | (y & z)

	var setBit byte
	if f == x {
		setBit = a.r1[SizeReg1-13-1] ^ a.r1[SizeReg1-16-1] ^ a.r1[SizeReg1-17-1] ^ a.r1[SizeReg1-18-1]
		a.r1 = rotateLeft(a.r1)
		a.r1[len(a.r1)-1] = setBit
	}

	if f == y {
		setBit = a.r2[SizeReg2-20-1] ^ a.r2[SizeReg2-21-1]
		a.r2 = rotateLeft(a.r2)
		a.r2[len(a.r2)-1] = setBit
	}

	if f == z {
		setBit = a.r3[SizeReg3-7-1] ^ a.r3[SizeReg3-20-1] ^ a.r3[SizeReg3-21-1] ^ a.r3[SizeReg3-22-1]
		a.r3 = rotateLeft(a.r3)
		a.r3[len(a.r3)-1] = setBit
	}
}

func rotateLeft(arr []byte) []byte {
	length := len(arr)
	if length == 0 {
		return nil
	}

	temp := make([]byte, length)

	//toDelete := arr[0]
	for i := 1; i < length; i++ {
		temp[(i - 1)] = arr[i]
	}
	//temp[length-1] = toDelete

	return temp
}

func (a *A51) initRegsMethod2(key []byte) {
	a.r1 = make([]byte, SizeReg1)
	a.r2 = make([]byte, SizeReg2)
	a.r3 = make([]byte, SizeReg3)

	var bit1, bit2, bit3 byte
	for i := 0; i < 64; i++ {
		a.r1[SizeReg1-1] = a.r1[SizeReg1-1] ^ a.r1[SizeReg1-2] ^ a.r1[SizeReg1-3] ^ key[SizeKey-i-1]
		a.r2[SizeReg2-1] = a.r2[SizeReg2-1] ^ a.r2[SizeReg2-2] ^ a.r2[SizeReg2-3] ^ key[SizeKey-i-1]
		a.r3[SizeReg3-1] = a.r3[SizeReg3-1] ^ a.r3[SizeReg3-2] ^ a.r3[SizeReg3-3] ^ key[SizeKey-i-1]

		a.shiftRegsWithoutSyncBit(bit1, bit2, bit3)
	}

	for i := 0; i < 223; i++ {
		a.shiftRegsWithSyncBit()
	}

	a.output()
}

func (a *A51) output() {
	for i := 0; i < 64; i++ {
		fmt.Printf("%3d: %b\n", i, a.r1)
		fmt.Printf("%3d: %b\n", i, a.r2)
		fmt.Printf("%3d: %b\n", i, a.r3)
	}
}
