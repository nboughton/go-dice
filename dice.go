package dice

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Dice represents a set of 1 type of dice, i.e: 3d20 OR 2d4 OR 1d6
// Types are not mixed as it is unnecessary
type Dice struct {
	Number, Sides int
	r             *rand.Rand
}

// NewDice returns a new Dice collection where
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

// Bag is a collection of different types of Dice, i.e: 3d20 AND 2d4 AND 1d6
type Bag []*Dice

// NewBag returns a new Bag object
func NewBag() Bag {
	return Bag{}
}

// Add puts more dice in the bag
func (b Bag) Add(d *Dice) {
	b = append(b, d)
}

// Roll returns aggregate rolls of all dice in the bag
func (b Bag) Roll() int {
	t := 0

	for _, d := range b {
		t += d.Roll()
	}

	return t
}

// String satisfies the Stringer interface for Bags
func (b Bag) String() string {
	v := make([]string, len(b))

	for i, d := range b {
		v[i] = fmt.Sprint(d)
	}

	return strings.Join(v, ", ")
}
