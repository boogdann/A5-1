package app

import (
	a51 "2/internal/a51/v2"
	"2/internal/ciphering"
	"2/internal/exel"
	"2/internal/files"
	"fmt"
	"io"
	"os"
	"strings"
)

type App struct {
	method   int
	filename string

	key uint64

	data    []byte
	keyBits []byte
	cipher  []byte
}

func New(method int, filename string, key uint64) *App {
	return &App{
		method:   method,
		filename: filename,
		key:      key,
	}
}

func (a *App) Run() ([]byte, []byte, []byte) {
	a5 := a51.New()

	err := a5.InitRegs(a.method, uint64ToBytesBits(a.key))
	if err != nil {
		panic(err)
	}

	data, err := Data(a.filename)
	if err != nil {
		panic(err)
	}

	crypt := ciphering.New(a5)
	cryptData, ke := crypt.Encrypt(data)

	a.data = data
	a.keyBits = ke
	a.cipher = cryptData

	return data, ke, cryptData
}

func (a *App) Save(pathTmplt, pathExelTmplt string) error {
	name := a.filename
	prevName := a.filename
	found := true
	for found {
		_, name, found = strings.Cut(name, "/")
		if !found {
			name = prevName
		}
		prevName = name
	}

	name, _, found = strings.Cut(name, ".")
	if !found {
		name = a.filename
	}

	err := files.Save(fmt.Sprintf(fmt.Sprintf("%s.txt", pathTmplt), name, a.method), a.cipher)
	if err != nil {
		return err
	}

	dataFreq := calcFrequency(getRawData(a.data))
	cipherFreq := calcFrequency(getRawData(a.cipher))

	exelFile := exel.New(128, fmt.Sprintf(fmt.Sprintf("%s.xlsx", pathExelTmplt), name, a.method))
	err = exelFile.Save(a.data, a.keyBits, a.cipher, dataFreq, cipherFreq)
	if err != nil {
		return err
	}

	return nil
}

func uint64ToBytesBits(num uint64) []byte {
	byteArr := make([]byte, 0, 64)
	for i := 63; i >= 0; i-- {
		byteArr = append(byteArr, byte((num>>i)&1))
	}
	return byteArr
}

func Data(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	newData := make([]byte, 0, len(data)*8)
	for _, num := range data {
		for i := 7; i >= 0; i-- {
			newData = append(newData, (num>>i)&1)
		}
	}

	return newData, nil
}

func calcFrequency(data []byte) []uint64 {
	freq := make([]uint64, 256)
	for i := 0; i < len(data); i++ {
		freq[data[i]]++
	}
	return freq
}

func getRawData(data []byte) []byte {
	saveData := make([]byte, len(data)/8)
	for i, bit := range data {
		byteIndex := i / 8
		bitIndex := (8 - 1) - uint(i%8)

		if bit == 1 {
			saveData[byteIndex] |= 1 << bitIndex
		}
	}

	return saveData
}
