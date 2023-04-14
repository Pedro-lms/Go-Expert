package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectedAmount := 5.0

	result := CalculateTax(amount)

	if result != expectedAmount {
		t.Errorf("Expected result %f, but got %f", expectedAmount, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expectedAmount float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{25000.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expectedAmount {
			t.Errorf("expected result %f, but got %f", item.expectedAmount, result)
		}

	}
}
