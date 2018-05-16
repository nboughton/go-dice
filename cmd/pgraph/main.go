package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/nboughton/go-dice"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var defaultrolls = 1000000

func main() {
	d := flag.String("d", "2d6", "Dice set to test. Can be a single value (2d10) or multiple values delineated by commas (2d4,3d10...)")
	r := flag.Int("r", defaultrolls, "Set number of rolls to try")
	flag.Parse()

	bag, err := dice.NewBag(strings.Split(*d, ",")...)
	if err != nil {
		log.Fatal(err)
	}

	// New plot
	pl, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	pl.Title.Text = "Distribution For " + bag.String()
	pl.X.Label.Text = "Rolled"
	pl.X.Min = float64(bag.Min())
	pl.X.Max = float64(bag.Max())

	pl.Y.Label.Text = "Probability (%)"
	pl.Y.Min = 0
	pl.Y.Max = 50

	pl.X.Tick.Marker = customTicks{}
	pl.Y.Tick.Marker = customTicks{}
	pl.Add(plotter.NewGrid())

	// Generate plot data
	l, err := plotter.NewLine(lineData(bag, *r))
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)

	// Add plot data
	pl.Add(l)

	// Save to png
	if err := pl.Save(20*vg.Centimeter, 15*vg.Centimeter, fmt.Sprintf("%s.png", *d)); err != nil {
		log.Fatal(err)
	}
}

func lineData(bag *dice.Bag, rolls int) plotter.XYs {
	var (
		min  = bag.Min()
		max  = bag.Max()
		xLen = max - min + 1
		xy   = make(plotter.XYs, xLen)
	)

	for i := 0; i < rolls; i++ {
		t, _ := bag.Roll()
		xy[t-min].Y++
	}

	for i := range xy {
		xy[i].X = float64(i + min)
		xy[i].Y = xy[i].Y / float64(rolls) * 100
	}

	return xy
}

type customTicks struct{}

func (customTicks) Ticks(min, max float64) []plot.Tick {
	var tks []plot.Tick

	for i := 0.; i < max; i++ {
		t := plot.Tick{Value: float64(i + 1)}

		switch {
		case max > 20 && max < 50:
			t.Label = label(i, 2)
		case max >= 50 && max < 100:
			t.Label = label(i, 5)
		case max >= 100:
			t.Label = label(i, int(max/4))
		default:
			t.Label = label(i, 1)
		}

		tks = append(tks, t)
	}

	return tks
}

func label(i float64, mod int) string {
	if int(i+1)%mod == 0 {
		return fmt.Sprintf("%d", int(i+1))
	}

	return ""
}
