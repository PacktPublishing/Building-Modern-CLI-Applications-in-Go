package widgets

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/barchart"
	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/gauge"
)

func BarChart() error {
	barChart, err := barchart.New()
	if err != nil {
		return err
	}
	values := []int{20, 40, 60, 80, 100}
	max := 100
	if err := barChart.Values(values, max); err != nil {
		return err
	}
	return nil
}

func Donut() error {
	greenDonut, err := donut.New(
		donut.CellOpts(cell.FgColor(cell.ColorGreen)),
		donut.Label("Green", cell.FgColor(cell.ColorGreen)),
	)
	if err != nil {
		return err
	}
	greenDonut.Percent(75)
	return nil
}

func Gauge() error {
	progressGauge, err := gauge.New(
		gauge.Height(1),
		gauge.Border(linestyle.Light),
		gauge.BorderTitle("Percentage progress"),
	)
	if err != nil {
		return err
	}
	progressGauge.Percent(75)
	return nil
}
