package main

import (
	"fmt"
	"time"

	ui "github.com/gizak/termui"
)

var (
	colorIndex = 0
	colors     = []ui.Attribute{
		ui.ColorGreen,
		ui.ColorYellow,
		ui.ColorMagenta,
		ui.ColorRed,
		ui.ColorCyan,
	}

	reverseColor = map[ui.Attribute]string{
		ui.ColorGreen:   "green",
		ui.ColorYellow:  "yellow",
		ui.ColorMagenta: "magenta",
		ui.ColorRed:     "red",
		ui.ColorCyan:    "cyan",
	}
)

type chartData struct {
	Capacity []float64
}

// func updateTerm() {
// 	tm.Clear()
// 	tm.MoveCursor(1, 1)

// 	chart := tm.NewLineChart(100, 20)
// 	data := new(tm.DataTable)
// 	rows := make([][]float64, 0)
// 	data.AddColumn("Index")
// 	for i := range mc {
// 		data.AddColumn(mc[i].Metric)
// 	}
// 	smallest := *maxSeries * 100
// 	for i := range mc {
// 		rows = append(rows, mc[i].Value)
// 		if len(mc[i].Value) <= smallest {
// 			smallest = len(mc[i].Value)
// 		}
// 	}

// 	// add empty rows
// 	for i := 0; i < *maxSeries-smallest; i++ {
// 		emptyrow := make([]float64, 0)
// 		emptyrow = append(emptyrow, float64(i))
// 		for range mc {
// 			emptyrow = append(emptyrow, 0.0)
// 		}
// 		data.AddRow(emptyrow...)
// 	}

// 	// fill in with rows we have data for
// 	shiftedIndex := *maxSeries - smallest
// 	for i := 0; i < smallest; i++ {
// 		combinedRow := make([]float64, 0)
// 		combinedRow = append(combinedRow, float64(shiftedIndex))
// 		shiftedIndex++
// 		for x := range rows {
// 			combinedRow = append(combinedRow, rows[x][i])
// 		}
// 		data.AddRow(combinedRow...)
// 	}

// 	tm.Println(chart.Draw(data))
// 	tm.Flush()
// }

func rotateColor() ui.Attribute {
	colorIndex++
	if colorIndex >= len(colors) {
		colorIndex = 0
	}
	return colors[colorIndex]
}

func updateTerm() {

	lc0 := ui.NewLineChart()
	lc0.BorderLabel = "Monitoring Firehose"
	lc0.Width = 80
	lc0.Height = 15
	lc0.X = 0
	lc0.Y = 0
	lc0.AxesColor = ui.ColorBlue
	lc0.LineColor["first"] = ui.ColorGreen | ui.AttrBold

	rows := make(map[string][]float64)
	smallest := *maxSeries * 100
	for i := range mc {
		if len(mc[i].Value) <= smallest {
			smallest = len(mc[i].Value)
		}
	}
	for i := range mc {
		key := fmt.Sprintf("%s/%s", mc[i].Job, mc[i].Index)
		rows[key] = mc[i].Value
	}
	lc0.Data = rows
	legend := ""
	for k := range rows { // TODO  WE NDEED TO SORT KEYS
		lc0.LineColor[k] = rotateColor()
		legend += fmt.Sprintf("%s = %s\n", k, reverseColor[lc0.LineColor[k]])
	}
	colorIndex = 0

	p2 := ui.NewParagraph(legend)
	p2.Height = 5
	p2.Width = 80
	p2.Y = 16
	p2.BorderLabel = "legend"
	p2.BorderFg = ui.ColorYellow

	ui.Render(lc0, p2)
}

func loopTerm() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				ui.Close()
				logger.Fatal("Quiting..")
			}
		default:
			time.Sleep(1 * time.Second)
			updateTerm()
		}

	}
}
