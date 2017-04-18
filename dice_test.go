package dice

import (
	"strings"
	"testing"
)

var (
	testDice    = "1d20"
	testBag     = []string{"1d20", "4d4", "6d6"}
	testBagMax  = 72
	testBagMin  = 11
	rollTests   = 10000
	expectedStr = "Expected [%s], got [%s]\n"
	otbStr      = "%s is out of bounds: %d\n"
)

func TestNewDice(t *testing.T) {
	d, _ := NewDice(testDice)
	if d.String() != testDice {
		t.Fatalf(expectedStr, testDice, d.String())
	}
}

func TestDiceAdd(t *testing.T) {
	d, _ := NewDice("3d20")
	e := "4d20"
	d.Add(testDice)
	if d.String() != e {
		t.Fatalf(expectedStr, e, d.String())
	}

	d.Add("2d6")
	if d.String() != e {
		t.Fatalf(expectedStr, e, d.String())
	}
}

func TestDiceRemove(t *testing.T) {
	d, _ := NewDice("3d20")
	e := "2d20"
	d.Remove(testDice)
	if d.String() != e {
		t.Fatalf(expectedStr, e, d.String())
	}

	d.Remove("2d6")
	if d.String() != e {
		t.Fatalf(expectedStr, e, d.String())
	}
}

func TestDiceRoll(t *testing.T) {
	d, _ := NewDice(testDice)
	for i := 0; i < rollTests; i++ {
		r := d.Roll()
		if r < 1 || r > d.sides {
			t.Fatalf(otbStr, testDice, r)
		}
	}
}

func TestNewBag(t *testing.T) {
	b, _ := NewBag(testBag...)
	e := strings.Join(testBag, ", ")
	if b.String() != e {
		t.Fatalf(expectedStr, e, b)
	}
}

func TestNewBagMulti(t *testing.T) {
	b, _ := NewBag("1d20", "3d20", "8d10")
	e := "4d20, 8d10"
	if b.String() != e {
		t.Fatalf(expectedStr, e, b.String())
	}
}

func TestBagAdd(t *testing.T) {
	b, _ := NewBag(testBag...)
	e := "2d20, 4d4, 6d6"
	b.Add(testDice)
	if b.String() != e {
		t.Fatalf(expectedStr, e, b.String())
	}

	d, e := "3d10", e+", 3d10"
	b.Add(d)
	if b.String() != e {
		t.Fatalf(expectedStr, e, b.String())
	}
}

func TestBagRemove(t *testing.T) {
	b, _ := NewBag(testBag...)
	e := "0d20, 4d4, 6d6"
	b.Remove("2d20")
	if b.String() != e {
		t.Fatalf(expectedStr, e, b.String())
	}

	d, e := "2d4", "0d20, 2d4, 6d6"
	b.Remove(d)
	if b.String() != e {
		t.Fatalf(expectedStr, e, b.String())
	}
}

func TestBagRoll(t *testing.T) {
	b, _ := NewBag(testBag...)
	for i := 0; i < rollTests; i++ {
		r := b.Roll()
		if r < testBagMin || r > testBagMax {
			t.Fatalf(otbStr, b, r)
		}
	}
}
