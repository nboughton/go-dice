package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"gonum.org/v1/plot/vg"

	"github.com/nboughton/go-dice"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

var defaultPrecision = 1000000

func main() {
	d := flag.String("d", "2d6", "Dice set to test. Can be a single value (2d10) or multiple values delineated by commas (2d4,3d10...)")
	p := flag.String("p", "high", "Set precision (high, medium, low). Higher precision performs more tests and thus takes longer")
	flag.Parse()

	precision := defaultPrecision

	switch *p {
	case "medium":
		precision = 100000
	case "low":
		precision = 10000
	}

	bag, err := dice.NewBag(strings.Split(*d, ",")...)
	if err != nil {
		log.Fatal(err)
	}

	// New plot
	pl, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	pl.Title.Text = "Probability Distribution For " + *d
	pl.X.Label.Text = "Rolled"
	pl.Y.Label.Text = "Probability (%)"
	pl.Add(plotter.NewGrid())

	// Generate plot data
	l, err := plotter.NewLine(lineData(bag, precision))
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)

	// Add plot data
	pl.Add(l)
	//pl.Legend.Add(fmt.Sprintf("For %s", *d))

	// Save to png
	if err := pl.Save(15*vg.Centimeter, 15*vg.Centimeter, fmt.Sprintf("%s.png", *d)); err != nil {
		log.Fatal(err)
	}
}

func runTests(bag *dice.Bag, tests, wants int) float64 {
	c := 0
	for i := 0; i < tests; i++ {
		t, _ := bag.Roll()
		if t == wants {
			c++
		}
	}

	return float64(c) / float64(tests) * 100
}

func lineData(bag *dice.Bag, precision int) plotter.XYs {
	pts := make(plotter.XYs, bag.Max()-bag.Min()+1)
	for i := range pts {
		pts[i].X = float64(i + bag.Min())
		pts[i].Y = runTests(bag, precision, i+bag.Min())
	}
	return pts
}
