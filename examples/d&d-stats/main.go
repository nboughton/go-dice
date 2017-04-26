package main

import (
	"fmt"
	"sort"

	"github.com/nboughton/go-dice"
)

func main() {
	// Run once for each stat
	for i := 0; i < 6; i++ {
		n := []int{}
		d, _ := dice.NewDice("1d6")

		// roll 4 times
		for j := 0; j < 4; j++ {
			n = append(n, d.Roll())
		}

		sort.Ints(n)
		fmt.Printf("Dropped: [%d], Kept: %v = %d\n", n[0], n[1:], sum(n[1:]))
	}
}

func sum(n []int) int {
	t := 0
	for _, v := range n {
		t += v
	}
	return t
}
