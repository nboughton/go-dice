package dice

import (
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
	b, e := NewBag(testBag...).String(), strings.Join(testBag, ", ")
	if b != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b)
	}
}

func TestNewBagMulti(t *testing.T) {
	b, e := NewBag("1d20", "3d20", "8d10"), "4d20, 8d10"
	if b.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b.String())
	}
}

func TestBagAdd(t *testing.T) {
	b, e := NewBag(testBag...), "2d20, 4d4, 6d6"
	b.Add(testDice)
	if b.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b.String())
	}

	d, e := "3d10", e+", 3d10"
	b.Add(d)
	if b.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b.String())
	}
}

func TestBagRemove(t *testing.T) {
	b, e := NewBag(testBag...), "0d20, 4d4, 6d6"
	b.Remove("2d20")
	if b.String() != e {
		t.Fatalf("Expected [%s], got [%s]\n", e, b.String())
	}

	d, e := "2d4", "0d20, 2d4, 6d6"
	b.Remove(d)
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
