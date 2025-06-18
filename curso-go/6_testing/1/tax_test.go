package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500
	exptected := 5.0

	result := CalculateTax(float64(amount))

	if result != exptected {
		t.Errorf("Expected %f, but got %f", exptected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500, 5.0},
		{1000, 10.0},
		{1500, 10.0},
		{999.99, 5.0},
		{100.00, 5.0},
		{0.00, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f, but got %f for amount %f", item.expected, result, item.amount)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64){
		result := CalculateTax(amount)
		if amount <= 0 && result != 0.0{
			t.Errorf("Received %f but expected 0", result)
		}
		if amount > 20000 && result != 20{
			t.Errorf("Received %f but expected 20", result)
		}
	})
}
