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
	filename      = "file.txt"
	saveFilename1 = "file_s.txt"
	saveFilename2 = "file_s.txt"
	dataFilename  = "data.xlsx"
	key           = uint64(12345678987654)
)

var (
	rawData []byte
)

func Run() {
	a5 := a51.New()

	fmt.Printf("KEY: %064b\n", key)

	err := a5.InitRegs(a51.Method2, uint64ToBytesBits(key))
	if err != nil {
		panic(err)
	}

	data, err := Data(filename)
	if err != nil {
		panic(err)
	}

	crypt := ciphering.New(a5)
	cryptData, ke := crypt.Encrypt(data)

	dataFreq := calcFrequency(rawData)

	if err := CreateExelFile(data, ke, cryptData, dataFreq, dataFreq); err != nil {
		panic(err)
	}

	if err := Save(cryptData); err != nil {
		panic(err)
	}
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

	rawData = data

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

func CreateExelFile(data []byte, key []byte, cryptData []byte, dataFreq, cryptFreq []uint64) (err error) {
	f := excelize.NewFile()
	defer func() {
		err = f.Close()
	}()

	f.NewSheet("Frequency Histogram")

	f.SetCellValue("Frequency Histogram", "A1", "ASCII Code")
	f.SetCellValue("Frequency Histogram", "B1", "Frequency")

	// Заполняем таблицу данными
	for i, count := range dataFreq {
		_ = count
		row := fmt.Sprintf("%d", i)
		f.SetCellValue("Frequency Histogram", fmt.Sprintf("A%d", i+2), row)
		f.SetCellValue("Frequency Histogram", fmt.Sprintf("B%d", i+2), 100)

	}

	// Создаем диаграмму
	err = f.AddChart("Frequency Histogram", "C1", &excelize.Chart{
		Type: excelize.Bar,
		Series: []excelize.ChartSeries{
			{
				Name:       "Frequency Histogram!$A$2",
				Categories: "Frequency Histogram!$A$1:!$257",
				Values:     "Frequency Histogram!$B$1:$B$257",
			},
			{
				Name:       "Frequency Histogram!$A$2",
				Categories: "Frequency Histogram!$A$1:!$257",
				Values:     "Frequency Histogram!$B$1:$B$257",
			},
		},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//err = f.SetColWidth("Sheet1", "A", "C", 5)
	//if err != nil {
	//	return
	//}
	//err = f.SetRowOutlineLevel("Sheet1", 2, 1)
	//
	//f.SetCellValue("Sheet1", "A1", "Data")
	//f.SetCellValue("Sheet1", "B1", "Key")
	//f.SetCellValue("Sheet1", "C1", "CryptData")
	//
	//size := min(sizeExel, len(data), len(cryptData), len(key))
	//for i := 0; i < size; i++ {
	//	f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), data[i])
	//	f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), key[i])
	//	f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), cryptData[i])
	//}

	err = f.SaveAs(dataFilename)
	return
}

func calcFrequency(data []byte) []uint64 {
	freq := make([]uint64, 256)
	for i := 0; i < len(data); i++ {
		freq[data[i]]++
	}
	return freq
}
