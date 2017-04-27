package main

import (
	"fmt"
	"strings"

	"github.com/nboughton/go-dice"
)

// Here's a little example that rolls Fate dice, used in the Fate family of tabeltop RPGS

var (
	fate = map[int]int{1: -1, 2: -1, 3: 0, 4: 0, 5: 1, 6: 1}
	sym  = map[int]string{-1: "[-]", 0: "[ ]", 1: "[+]"}
)

func main() {
	t, a := 0, []string{}
	d, _ := dice.NewDice("4d6")
	_, s := d.Roll()

	for _, n := range s {
		a = append(a, sym[fate[n]])
		t += fate[n]
	}

	fmt.Printf("%s = %d\n", strings.Join(a, ", "), t)
}
