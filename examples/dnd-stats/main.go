package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/nboughton/go-dice"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 1, ' ', 0)

	// Run once for each stat
	for i := 0; i < 6; i++ {
		d, _ := dice.NewDice("4d6")
		_, n := d.Roll()
		sort.Ints(n)
		fmt.Fprintf(w, "%d\t%v\n", sum(n[1:]), n)
	}

	w.Flush()
}

func sum(n []int) int {
	t := 0
	for _, v := range n {
		t += v
	}
	return t
}
