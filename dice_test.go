package dice

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testDice   = "1d20"
	testBag    = []string{"1d20", "4d4", "6d6"}
	testBagMax = 72
	testBagMin = 11
)

func TestNewDice(t *testing.T) {
	if NewDice(testDice).String() != testDice {
		t.Fatal("Expected ", testDice)
	}
}

func TestDiceRoll(t *testing.T) {
	d := NewDice(testDice)
	for i := 0; i < 10000; i++ {
		r := d.Roll()
		if r < 1 || r > d.Sides {
			t.Fatalf("Dice (%s) rolled out of bounds: %d\n", testDice, r)
		}
	}
}

func TestNewBag(t *testing.T) {
	b, v := NewBag(testBag...).String(), strings.Join(testBag, ", ")
	if b != v {
		t.Fatalf("Expected [%s], got [%s]\n", v, b)
	}
}

func TestBagAdd(t *testing.T) {
	b := NewBag(testBag...)
	b.Add(testDice)
	tBStr := "2d20, 4d4, 6d6"
	if b.String() != tBStr {
		t.Fatalf("Expected [%s], got [%s]\n", tBStr, b.String())
	}

	d := "3d10"
	b.Add(d)
	tBStr = fmt.Sprintf("%s, %s", tBStr, d)
	if b.String() != tBStr {
		t.Fatalf("Expected [%s], got [%s]\n", tBStr, b.String())
	}
}

func TestBagRemove(t *testing.T) {
	b := NewBag(testBag...)
	b.Remove("2d20")
	tBStr := "0d20, 4d4, 6d6"
	if b.String() != tBStr {
		t.Fatalf("Expected [%s], got [%s]\n", tBStr, b.String())
	}

	d := "2d4"
	b.Remove(d)
	tBStr = "0d20, 2d4, 6d6"
	if b.String() != tBStr {
		t.Fatalf("Expected [%s], got [%s]\n", tBStr, b.String())
	}
}

func TestBagRoll(t *testing.T) {
	b := NewBag(testBag...)
	for i := 0; i < 10000; i++ {
		r := b.Roll()
		if r < testBagMin || r > testBagMax {
			t.Fatalf("Bag roll [%d] out of bounds", r)
		}
	}
}
