package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/nboughton/go-dice"
)

var tw = tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must specify dice to roll, i.e 1d6 or 1d6 2d4 3d20 etc")
	}

	bag, err := dice.NewBag(os.Args[1:]...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	total, ds := bag.Roll()

	str := []string{fmt.Sprintf("Total\t:\t%d", total)}
	for die, num := range ds {
		str = append(str, fmt.Sprintf("%s\t:\t%v", die, num))
	}

	fmt.Fprintln(tw, strings.Join(str, "\n"))
	tw.Flush()
}
