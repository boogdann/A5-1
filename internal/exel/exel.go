package exel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type File struct {
	sizeExel     int
	exelFilename string
}

func New(sizeExel int, exelFilename string) *File {
	return &File{
		sizeExel:     sizeExel,
		exelFilename: exelFilename,
	}
}

func (e *File) Save(data []byte, key []byte, cryptData []byte, dataFreq, cryptFreq []uint64) (err error) {
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

	size := min(e.sizeExel, len(data), len(cryptData), len(key))
	for i := 0; i < size; i++ {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), data[i])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), key[i])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), cryptData[i])
	}

	err = f.SaveAs(e.exelFilename)
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
