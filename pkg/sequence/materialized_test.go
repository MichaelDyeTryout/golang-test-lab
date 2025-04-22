package sequence

import (
	"math"
	"testing"
)

func TestCountingNumbersEmpty(t *testing.T) {

	qtys := []int8{0, -1, -2}

	for _, qty := range qtys {
		seq, err := CountingNumbers(qty)

		if err != nil {
			t.Errorf("Error should be nil for qty %d but got %v", qty, err)
			return // remember that golang t.Errorf() continues execution
		}

		if seq == nil {
			t.Errorf("Counting numbers slice nil for given input qty %d", qty)
			return // remember that golang t.Errorf() continues execution
		}

		if len(*seq) != 0 {
			t.Errorf("Requested qty %d counting numbers and expected 0 but got %d", qty, len(*seq))
		}
	}
}

func TestCountingNumbersBoundaryLow(t *testing.T) {
	// low boundary for counting numbers is 1

	seq, err := CountingNumbers(1)
	head := (*seq)[0]

	if head != 1 {
		t.Errorf("Requested 1 counting number and expected value 1 but got %d", head)
	}
	if err != nil {
		t.Errorf("Error should be nil but got %v", err)
	}
}

func TestCountingNumbersBoundaryHigh(t *testing.T) {
	// high boundary for counting numbers is math.MaxInt

	seq, err := CountingNumbers(math.MaxInt8)
	tail := (*seq)[len(*seq)-1]

	if tail != math.MaxInt8 {
		t.Errorf("Requested %d counting numbers but got %d", math.MaxInt8, tail)
	}
	if err != nil {
		t.Errorf("Error should be nil but got %v", err)
	}
}

func TestCountingNumbersCorrectQuantity(t *testing.T) {

	qty := int8(2)

	seq, err := CountingNumbers(qty)
	if len(*seq) != 2 {
		t.Errorf("Requested %d counting numbers but got %d", qty, len(*seq))
	}
	if err != nil {
		t.Errorf("Error should be nil but got %v", err)
	}
}

func TestIntegerRangeInclusiveEmpty(t *testing.T) {
	begin := int8(0)
	end := int8(-2)

	seq, err := IntegerRange(begin, end)

	if err != nil {
		t.Errorf("Error should be nil for (%d, %d) but got %v", begin, end, err)
		return // remember that golang t.Errorf() continues execution
	}

	if seq == nil {
		t.Errorf("Integer range slice nil for given input (%d, %d)", begin, end)
		return // remember that golang t.Errorf() continues execution
	}

	if len(*seq) != 0 {
		t.Errorf("Requested range (%d, %d) and got %d", begin, end, len(*seq))
	}
}

func TestRepeatBasic(t *testing.T) {

	qty := int8(2)

	seq, err := Repeat(qty, 2)

	if len(*seq) != 2 {
		t.Errorf("Requested %d repeated numbers but got %d", qty, len(*seq))
	}

	for ix := range *seq {
		if (*seq)[ix] != 2 {
			t.Errorf("Expected value 2 but got %d", (*seq)[ix])
		}
	}

	if err != nil {
		t.Errorf("Error should be nil but got %v", err)
	}
}
