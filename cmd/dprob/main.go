package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/nboughton/go-dice"
)

// This is an example utility I use to test the probability of results by dice combination

func main() {
	// Parse flags
	d := flag.String("d", "1d20,1d4", "Dice set as comma separated list [1d20,3d4...]")
	e := flag.Int("e", 2, "Expected result to test")
	t := flag.Int("t", 10000, "Number of tests to run")
	flag.Parse()

	// Expand *d into sets of test dice
	diceBag, err := dice.NewBag(strings.Split(*d, ",")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run Tests
	fmt.Printf("Probability of %v == %d: %f%% over %d tests\n", diceBag, *e, runTests(diceBag, *t, *e), *t)
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
