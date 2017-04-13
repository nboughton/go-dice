package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Dice represents a set of 1 type of dice, i.e: 3d20 OR 2d4 OR 1d6
type Dice struct {
	Number, Sides int
	r             *rand.Rand
}

// NewDice returns a new Dice set
func NewDice(number, sides int) *Dice {
	return &Dice{number, sides, rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// Roll all dice in set and return the aggregate result
func (d *Dice) Roll() int {
	t := 0
	for i := 0; i < d.Number; i++ {
		t += d.r.Intn(d.Sides) + 1
	}
	return t
}

// String satisfies the Stringer interface for Dice
func (d *Dice) String() string {
	return fmt.Sprintf("%dd%d", d.Number, d.Sides)
}

// Bag is a collection of different types of Dice, i.e: [3d20, 2d4, 1d6]
type Bag []*Dice

// NewBag returns a new Bag object. A bag can be created with a collection of
// dice specified in string form for convenience. I.e b := NewBag("2d20", "1d6", "8d8")
func NewBag(dice ...string) *Bag {
	b := &Bag{}

	//	if len(dice) > 0 {
	for _, a := range dice {
		var (
			ns   = strings.Split(a, "d")
			n, _ = strconv.Atoi(ns[0])
			s, _ = strconv.Atoi(ns[1])
		)

		*b = append(*b, NewDice(n, s))
	}
	//	}

	return b
}

// Add puts more dice in the bag
func (b *Bag) Add(d *Dice) {
	*b = append(*b, d)
}

// Roll returns aggregate rolls of all Dice in the bag
func (b *Bag) Roll() int {
	t := 0

	for _, d := range *b {
		t += d.Roll()
	}

	return t
}

// String satisfies the Stringer interface for Bags
func (b *Bag) String() string {
	v := make([]string, len(*b))

	for i, d := range *b {
		v[i] = fmt.Sprint(d)
	}

	return strings.Join(v, ", ")
}
