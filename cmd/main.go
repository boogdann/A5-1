package main

import "2/internal/app"

func main() {
	app.Run()
}

//func main() {
//	f := excelize.NewFile()
//	defer func() {
//		if err := f.Close(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	for idx, row := range [][]interface{}{
//		{nil, "Apple", "Orange", "Pear"},
//		{"Small", 2, 3, 3},
//		{"Normal", 5, 2, 4},
//		{"Large", 6, 7, 8},
//	} {
//		cell, err := excelize.CoordinatesToCellName(1, idx+1)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		f.SetSheetRow("Sheet1", cell, &row)
//	}
//	if err := f.AddChart("Sheet1", "E1", &excelize.Chart{
//		Type: excelize.Col3DClustered,
//		Series: []excelize.ChartSeries{
//			{
//				Name:       "Sheet1!$A$2",
//				Categories: "Sheet1!$B$1:$D$1",
//				Values:     "Sheet1!$B$2:$D$2",
//			},
//			{
//				Name:       "Sheet1!$A$3",
//				Categories: "Sheet1!$B$1:$D$1",
//				Values:     "Sheet1!$B$3:$D$3",
//			},
//			{
//				Name:       "Sheet1!$A$4",
//				Categories: "Sheet1!$B$1:$D$1",
//				Values:     "Sheet1!$B$4:$D$4",
//			},
//		},
//		Title: []excelize.RichTextRun{
//			{
//				Text: "Fruit 3D Clustered Column Chart",
//			},
//		},
//		Legend: excelize.ChartLegend{
//			ShowLegendKey: false,
//		},
//		PlotArea: excelize.ChartPlotArea{
//			ShowBubbleSize:  true,
//			ShowCatName:     false,
//			ShowLeaderLines: false,
//			ShowPercent:     true,
//			ShowSerName:     true,
//			ShowVal:         true,
//		},
//	}); err != nil {
//		fmt.Println(err)
//		return
//	}
//	// Save spreadsheet by the given path.
//	if err := f.SaveAs("Book1.xlsx"); err != nil {
//		fmt.Println(err)
//	}
//}
