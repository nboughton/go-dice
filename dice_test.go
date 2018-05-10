package dice

import (
	"strings"
	"testing"
)

var (
	testDice      = "1d20"
	testBag       = []string{"1d20", "4d4", "6d6"}
	testBagMax    = 72
	testBagMin    = 11
	rollTests     = 1000000
	expectedStr   = "Expected [%s], got [%s]\n"
	outOfBoundStr = "%s is out of bounds: %d\n"
	badAggSetStr  = "%s should contain %d elements, has %d\n"
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
	d.Add(1)
	if d.String() != e {
		t.Fatalf(expectedStr, e, d.String())
	}
}

func TestDiceRemove(t *testing.T) {
	d, _ := NewDice("3d20")
	e := "2d20"
	d.Remove(1)
	if d.String() != e {
		t.Fatalf(expectedStr, e, d.String())
	}
}

func TestDiceRoll(t *testing.T) {
	d, _ := NewDice(testDice)
	for i := 0; i < rollTests; i++ {
		r, s := d.Roll()
		// test individual roll
		if r < 1 || r > d.sides {
			t.Fatalf(outOfBoundStr, testDice, r)
		}
		// test set
		if len(s) != d.number {
			t.Fatalf(badAggSetStr, s, d.number, len(s))
		}
	}
}

func TestDiceMin(t *testing.T) {
	d, _ := NewDice(testDice)
	if d.Min() != 1 {
		t.Fatalf(expectedStr, 1, d.Min())
	}
}

func TestDiceMax(t *testing.T) {
	d, _ := NewDice(testDice)
	if d.Max() != 20 {
		t.Fatalf(expectedStr, 20, d.Max())
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
		r, s := b.Roll()
		// Check rolls
		if r < testBagMin || r > testBagMax {
			t.Fatalf(outOfBoundStr, b, r)
		}
		// Check set maps
		if len(s) != len(testBag) {
			t.Fatalf(badAggSetStr, s, len(testBag), len(s))
		}
	}
}

func TestBagMin(t *testing.T) {
	b, _ := NewBag(testBag...)
	if b.Min() != 11 {
		t.Fatalf(expectedStr, 11, b.Min())
	}
}

func TestBagMax(t *testing.T) {
	b, _ := NewBag(testBag...)
	if b.Max() != 72 {
		t.Fatalf(expectedStr, 72, b.Max())
	}
}
