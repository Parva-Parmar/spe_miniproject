package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	if got := Add(2.5, 3.5); got != 6.0 {
		t.Errorf("Add(2.5, 3.5) = %f; want 6.0", got)
	}
}

func TestSubtract(t *testing.T) {
	if got := Subtract(10.0, 4.0); got != 6.0 {
		t.Errorf("Subtract(10.0, 4.0) = %f; want 6.0", got)
	}
}

func TestMultiply(t *testing.T) {
	if got := Multiply(3.0, 4.0); got != 12.0 {
		t.Errorf("Multiply(3.0, 4.0) = %f; want 12.0", got)
	}
}

func TestDivide(t *testing.T) {
	if got := Divide(10.0, 2.0); got != 5.0 {
		t.Errorf("Divide(10.0, 2.0) = %f; want 5.0", got)
	}
}

func TestSquareRoot(t *testing.T) {
	if got := SquareRoot(16); got != 4 {
		t.Errorf("SquareRoot(16) = %f; want 4", got)
	}
}

func TestFactorial(t *testing.T) {
	if got := Factorial(5); got != 120 {
		t.Errorf("Factorial(5) = %d; want 120", got)
	}
	if got := Factorial(0); got != 1 {
		t.Errorf("Factorial(0) = %d; want 1", got)
	}
}

func TestNaturalLog(t *testing.T) {
	expected := math.Log(math.E)
	if got := NaturalLog(math.E); got != expected {
		t.Errorf("NaturalLog(e) = %f; want %f", got, expected)
	}
}

func TestPower(t *testing.T) {
	if got := Power(2, 3); got != 8 {
		t.Errorf("Power(2, 3) = %f; want 8", got)
	}
}
