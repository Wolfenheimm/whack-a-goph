package packages

import "testing"

func TestCalculator(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		calc := NewCalculator()

		calc.Add(5)
		if got := calc.Result(); got != 5 {
			t.Errorf("Add(5) = %v; want 5", got)
		}

		calc.Subtract(2)
		if got := calc.Result(); got != 3 {
			t.Errorf("Subtract(2) = %v; want 3", got)
		}

		calc.Multiply(4)
		if got := calc.Result(); got != 12 {
			t.Errorf("Multiply(4) = %v; want 12", got)
		}

		err := calc.Divide(3)
		if err != nil {
			t.Errorf("Divide(3) returned error: %v", err)
		}
		if got := calc.Result(); got != 4 {
			t.Errorf("Divide(3) = %v; want 4", got)
		}
	})

	t.Run("Division by Zero", func(t *testing.T) {
		calc := NewCalculator()
		calc.Add(10)
		err := calc.Divide(0)
		if err == nil {
			t.Error("Divide(0) did not return error")
		}
	})
}
