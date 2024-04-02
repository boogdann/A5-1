package main

import (
	a51 "2/internal/a51/v2"
	"2/internal/exel"
	"2/internal/nist/discrete"
	"2/internal/nist/freqblock"
	"2/internal/nist/frequency"
	"2/internal/nist/rank"
	"2/internal/nist/runs"
	"2/internal/nist/runsblock"
	"fmt"
	"sync"
)

const (
	blockSize    = 16
	m            = 2
	q            = 2
	filename     = "tests/textfile1.txt"
	saveFilename = "tests/cipher/textfile1.method_1.save.txt"
	exelPath1    = "tests/exel/textfile1.method_1.save.xlsx"
	exelPath2    = "tests/exel/textfile1.method_2.save.xlsx"
	key          = uint64(463468934894)
)

func main() {
	a := a51.New()
	maxSize := max(128, 38*m*q*8)
	a.InitRegs(a51.Method1, uint64ToBytesBits(key))
	ke := a.GenerateKeyStream(maxSize - 1)

	for _, t := range ke {
		fmt.Printf("%d", t)
	}

	fmt.Printf("\n Len: %d", len(ke))
	a = a51.New()
	a.InitRegs(a51.Method2, uint64ToBytesBits(key))
	ke2 := a.GenerateKeyStream(maxSize - 1)

	for _, t := range ke2 {
		fmt.Printf("%d", t)
	}

	tests(ke, exelPath1)
	tests(ke2, exelPath2)
}

func uint64ToBytesBits(num uint64) []byte {
	byteArr := make([]byte, 0, 64)
	for i := 63; i >= 0; i-- {
		byteArr = append(byteArr, byte((num>>i)&1))
	}
	return byteArr
}

func tests(bits []byte, exelPath string) {

	var wg sync.WaitGroup
	wg.Add(6)

	func() {
		test := discrete.New(bits)
		got := test.Run()

		e := exel.New(0, exelPath)
		err := e.SaveTests(6, "DFT", got)
		_ = err
		wg.Done()
	}()

	func() {
		test1 := freqblock.New(bits, blockSize)
		got1 := test1.Run()

		e1 := exel.New(0, exelPath)
		e1.SaveTests(2, fmt.Sprintf("BlockFrequency (block - %d)", blockSize), got1)
		wg.Done()
	}()

	func() {
		test3 := frequency.New(bits)
		got3 := test3.Run()

		e3 := exel.New(0, exelPath)
		e3.SaveTests(1, "Frequency", got3)
		wg.Done()
	}()

	func() {
		test4 := rank.New(bits, m, q)
		got4, _ := test4.Run()

		e4 := exel.New(0, exelPath)
		e4.SaveTests(5, fmt.Sprintf("Rank (m - %d, q - %d)", m, q), got4)
		wg.Done()
	}()

	func() {
		test5 := runs.New(bits)
		got5 := test5.Run()

		e5 := exel.New(0, exelPath)
		e5.SaveTests(3, "Rans", got5)
		wg.Done()
	}()

	func() {
		test6 := runsblock.New(bits)
		got6, _ := test6.Run()

		e6 := exel.New(0, exelPath)
		e6.SaveTests(4, "Runsblock", got6)
		wg.Done()
	}()

	wg.Wait()
}
