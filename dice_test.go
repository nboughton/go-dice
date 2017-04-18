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
	rollTests  = 10000
)

func TestNewDice(t *testing.T) {
	if NewDice(testDice).String() != testDice {
		t.Fatal("Expected ", testDice)
	}
}

func TestDiceAdd(t *testing.T) {
	d, e := NewDice("3d20"), "4d20"
	d.Add(testDice)
	if d.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, d.String())
	}

	d.Add("2d6")
	if d.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, d.String())
	}
}

func TestDiceRemove(t *testing.T) {
	d, e := NewDice("3d20"), "2d20"
	d.Remove(testDice)
	if d.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, d.String())
	}

	d.Remove("2d6")
	if d.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, d.String())
	}
}

func TestDiceRoll(t *testing.T) {
	d := NewDice(testDice)
	for i := 0; i < rollTests; i++ {
		r := d.Roll()
		if r < 1 || r > d.sides {
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

func TestNewBagMulti(t *testing.T) {
	b, e := NewBag("1d20", "3d20", "8d10"), "4d20, 8d10"
	if b.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b.String())
	}
}

func TestBagAdd(t *testing.T) {
	b := NewBag(testBag...)
	b.Add(testDice)
	e := "2d20, 4d4, 6d6"
	if b.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b.String())
	}

	d := "3d10"
	b.Add(d)
	e = fmt.Sprintf("%s, %s", e, d)
	if b.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b.String())
	}
}

func TestBagRemove(t *testing.T) {
	b := NewBag(testBag...)
	b.Remove("2d20")
	e := "0d20, 4d4, 6d6"
	if b.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b.String())
	}

	d := "2d4"
	b.Remove(d)
	e = "0d20, 2d4, 6d6"
	if b.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b.String())
	}
}

func TestBagRoll(t *testing.T) {
	b := NewBag(testBag...)
	for i := 0; i < rollTests; i++ {
		r := b.Roll()
		if r < testBagMin || r > testBagMax {
			t.Fatalf("Bag roll [%d] out of bounds", r)
		}
	}
}
