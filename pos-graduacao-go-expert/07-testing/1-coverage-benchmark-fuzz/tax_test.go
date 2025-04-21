package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	// pelo método nativo
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expected {
			t.Errorf("Expected %f, got %f", item.expected, result)
		}
	}
}

// Verificar performance da função
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// tela de fuzzing, muito útil para encontrar problemas
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-200.0, -30.00, 0.0, 500.0, 1000.0, 1500.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount < 0 && result != 0.0 {
			t.Errorf("Expected %f, got %f", 0.0, result)
		}
		if amount >= 0 && amount < 1000 && result != 5.0 {
			t.Errorf("Expected %f, got %f", 5.0, result)
		}
		if amount > 20000 && result != 20.0 {
			t.Errorf("Expected %f, got %f", 20.0, result)
		}
	})
}
