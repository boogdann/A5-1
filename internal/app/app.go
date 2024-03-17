package app

import (
	a51 "2/internal/a51/v2"
	"2/internal/ciphering"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"os"
)

const (
	sizeExel      = 128
	filename      = "picture.png"
	saveFilename1 = "picture_save.png"
	saveFilename2 = "picture_save.txt"
	dataFilename  = "data.xlsx"
	key           = uint64(422672389234853)
)

func Run() {
	a5 := a51.New()

	fmt.Printf("KEY: %064b\n", key)

	err := a5.InitRegs(a51.Method1, uint64ToBytesBits(key))
	if err != nil {
		panic(err)
	}

	data, err := Data(filename)
	if err != nil {
		panic(err)
	}

	crypt := ciphering.New(a5)
	cryptData, ke := crypt.Encrypt(data)

	if err := CreateExelFile(data, ke, cryptData); err != nil {
		panic(err)
	}

	if err := Save(cryptData); err != nil {
		panic(err)
	}
}

func uint64ToBytesBits(num uint64) []byte {
	byteArr := make([]byte, 64)
	for i := 0; i < 64; i++ {
		byteArr[i] = byte((num >> i) & 1)
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
func Save(data []byte) error {
	file, err := os.Create(saveFilename1)
	if err != nil {
		return err
	}

	saveData := make([]byte, len(data)/8)
	for i, bit := range data {
		byteIndex := i / 8
		bitIndex := (8 - 1) - uint(i%8)

		if bit == 1 {
			saveData[byteIndex] |= 1 << bitIndex
		}
	}

	_, err = file.Write(saveData)

	file, err = os.Create(saveFilename2)
	if err != nil {
		return err
	}

	_, err = file.Write(saveData)

	return err
}

func CreateExelFile(data []byte, key []byte, cryptData []byte) (err error) {
	f := excelize.NewFile()
	defer func() {
		err = f.Close()
	}()

	err = f.SetColWidth("Sheet1", "A", "C", 5)
	if err != nil {
		return
	}
	err = f.SetRowOutlineLevel("Sheet1", 2, 1)

	f.SetCellValue("Sheet1", "A1", "Data")
	f.SetCellValue("Sheet1", "B1", "Key")
	f.SetCellValue("Sheet1", "C1", "CryptData")

	size := min(sizeExel, len(data), len(cryptData), len(key))
	for i := 0; i < size; i++ {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), data[i])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), key[i])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), cryptData[i])
	}

	err = f.SaveAs(dataFilename)
	return
}
