package app

import (
	a51 "2/internal/a51/v2"
	"2/internal/ciphering"
	"2/internal/nist/freqblock"
	"2/internal/nist/frequency"
	"2/internal/nist/runs"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"os"
)

const (
	sizeExel      = 128
	filename      = "ТИ_2.docx"
	saveFilename1 = "file_s.txt"
	saveFilename2 = "file_s.txt"
	dataFilename  = "data.xlsx"
	key           = uint64(81490833476438589)
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

	d := make([]byte, 0)
	d1 := make([]byte, 0)
	d = append(d, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1)
	d1 = append(d1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1)

	freqTest := frequency.New(d)
	a := freqTest.Run()
	fmt.Printf("A: %f\n", a)

	freqBlockTest := freqblock.New(d1, 3)
	b := freqBlockTest.Run()
	fmt.Printf("B: %f\n", b)

	runsTest := runs.New(d1)
	c := runsTest.Run()
	fmt.Printf("C: %f\n", c)

	dataFreq := calcFrequency(rawData)
	cryptRawData := getRawData(cryptData)
	cryptFreq := calcFrequency(cryptRawData)

	if err := CreateExelFile(data, ke, cryptData, dataFreq, cryptFreq); err != nil {
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

	add(f, dataFreq, "Frequency Histogram Data")
	add(f, cryptFreq, "Frequency Histogram CryptData")

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

func add(f *excelize.File, data []uint64, name string) {
	f.NewSheet(name)

	inter := make([][]interface{}, len(data))
	for i, count := range data {
		inter[i] = []interface{}{i, count}
	}

	series := make([]excelize.ChartSeries, 256)
	for idx, row := range inter {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetSheetRow(name, cell, &row)
		series[idx] = excelize.ChartSeries{
			Name:       fmt.Sprintf("'%s'!$A$%d", name, idx+1),
			Categories: fmt.Sprintf("'%s'!$B$%d:$D$%d", name, idx+1, idx+1),
			Values:     fmt.Sprintf("'%s'!$B$%d:$D$%d", name, idx+1, idx+1),
		}
	}

	tr := true
	if err := f.AddChart(name, "E1", &excelize.Chart{
		Format: excelize.GraphicOptions{
			AutoFit: false,
			Locked:  &tr,
			ScaleX:  3,
			ScaleY:  10,
		},
		Dimension: excelize.ChartDimension{
			Width:  500,
			Height: 200,
		},
		Type:   excelize.Bar,
		Series: series[:255],
		Title: []excelize.RichTextRun{
			{
				Text: "Frequency Histogram",
			},
		},
		Legend: excelize.ChartLegend{
			ShowLegendKey: false,
		},
		PlotArea: excelize.ChartPlotArea{
			ShowBubbleSize:  false,
			ShowCatName:     false,
			ShowLeaderLines: false,
			ShowPercent:     false,
			ShowSerName:     false,
			ShowVal:         true,
		},
	}); err != nil {
		fmt.Println(err)
		return
	}
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
