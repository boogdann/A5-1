package main

import (
	a51 "2/internal/a51/v2"
	"2/internal/app"
	"os"
)

const (
	//blockSize    = 32
	//m            = 32
	//q            = 32
	filename     = "tests/textfile1.txt"
	saveFilename = "tests/cipher/textfile1.method_1.save.txt"
	exelPath1    = "tests/exel/textfile1.method_1.save.xlsx"
	exelPath2    = "tests/exel/textfile1.method_2.save.xlsx"
	key          = uint64(123456789)
)

func main() {
	a := app.New(a51.Method2, filename, key)
	a.Run()

	if err := a.Save("tests/cipher/%s.method_%d.save",
		"tests/exel/%s.method_%d.save"); err != nil {
		panic(err)
	}
}

//func tests(bits []byte, exelPath string) {
//	test := discrete.New(bits)
//	got := test.Run()
//
//	e := exel.New(0, exelPath)
//	err := e.SaveTests(1, "Discrete", got)
//	_ = err
//
//	test1 := freqblock.New(bits, blockSize)
//	got1 := test1.Run()
//
//	e1 := exel.New(0, exelPath)
//	err = e1.SaveTests(2, fmt.Sprintf("BlockFrequency (block - %d)", blockSize), got1)
//
//	test3 := frequency.New(bits)
//	got3 := test3.Run()
//
//	e3 := exel.New(0, exelPath)
//	err = e3.SaveTests(3, "Frequency", got3)
//
//	test4 := rank.New(bits, m, q)
//	got4, _ := test4.Run()
//
//	e4 := exel.New(0, exelPath)
//	err = e4.SaveTests(4, fmt.Sprintf("Rank (m - %d, q - %d)", m, q), got4)
//
//	test5 := runs.New(bits)
//	got5 := test5.Run()
//
//	e5 := exel.New(0, exelPath)
//	err = e5.SaveTests(5, "Rank", got5)
//
//	test6 := runsblock.New(bits)
//	got6, _ := test6.Run()
//
//	e6 := exel.New(0, exelPath)
//	err = e6.SaveTests(6, "Runsblock", got6)
//}

func readTestData(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	newData := make([]byte, 0, len(data)*8)
	for _, num := range data {
		for i := 7; i >= 0; i-- {
			newData = append(newData, (num>>i)&1)
		}
	}

	return newData
}
