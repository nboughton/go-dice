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
	number, sides int
	r             *rand.Rand
}

// NewDice takes the common notation "nds" where n is the number of dice and s is the number of sides
// i.e 1d6 and returns a new Dice set
func NewDice(s string) *Dice {
	number, sides := strToVal(s)
	return &Dice{number, sides, rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// Roll all dice in set and return the aggregate result
func (d *Dice) Roll() int {
	t := 0

	for i := 0; i < d.number; i++ {
		t += d.r.Intn(d.sides) + 1
	}

	return t
}

// String satisfies the Stringer interface for Dice
func (d *Dice) String() string {
	return fmt.Sprintf("%dd%d", d.number, d.sides)
}

// Bag is a collection of different types of Dice, i.e: [3d20, 2d4, 1d6]
type Bag struct {
	dice []*Dice
}

// NewBag returns a new Bag object. A bag can be created with a collection of
// dice specified in string form for convenience. I.e b := NewBag("2d20", "1d6", "8d8")
func NewBag(dice ...string) *Bag {
	b := &Bag{}

	for _, a := range dice {
		b.Add(a)
	}

	return b
}

// Add puts more dice in the bag adding to existing sets where possible
func (b *Bag) Add(s string) {
	d := NewDice(s)

	// increment existing set if it exists
	for _, set := range b.dice {
		if set.sides == d.sides {
			set.number += d.number
			return
		}
	}

	// Otherwise add a new set
	b.dice = append(b.dice, d)
}

// Remove reduces the number of dice by the specified s string if s exists
func (b *Bag) Remove(s string) {
	number, sides := strToVal(s)

	// Remove specified dice from set
	for _, set := range b.dice {
		if set.sides == sides {
			// ensure no < 0 values
			if set.number-number < 0 {
				set.number = 0
			} else {
				set.number -= number
			}
			break
		}
	}
}

// Roll returns aggregate rolls of all Dice in the bag
func (b *Bag) Roll() int {
	t := 0

	for _, d := range b.dice {
		t += d.Roll()
	}

	return t
}

// String satisfies the Stringer interface for Bags
func (b *Bag) String() string {
	v := make([]string, len(b.dice))

	for i, d := range b.dice {
		v[i] = fmt.Sprint(d)
	}

	return strings.Join(v, ", ")
}

// returns int values for numbers, sides
func strToVal(a string) (number int, sides int) {
	ns := strings.Split(a, "d")
	number, _ = strconv.Atoi(ns[0])
	sides, _ = strconv.Atoi(ns[1])
	return number, sides
}
