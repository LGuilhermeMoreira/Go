package tax

import (
	"testing"
)

// TestCalculateTax is a unit test for the CalculateTax function.
func TestCalculateTax(t *testing.T) {
	amount := 600.0

	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// TestCalculateTaxBatch is another unit test for the CalculateTax function, using a table-driven test approach.
func TestCalculateTaxBatch(t *testing.T) {
	// Define a structure to hold test cases with input amount and expected result.
	type calcTax struct {
		amount, expect float64
	}

	// Create a table of test cases.
	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{750.0, 5.0},
		{0, 0},
	}

	// Iterate through the table and perform tests for each case.
	for _, item := range table {
		result := CalculateTax(item.amount)

		// Check if the result matches the expected value.
		if result != item.expect {
			t.Errorf("Expected %v, but got %v", item.expect, result)
		}
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}

	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Recived %v, but expected 0", result)
		}

		if amount >= 20000 && result != 20 {
			t.Errorf("Recived %v, but expected 20", result)
		}
	})
}
