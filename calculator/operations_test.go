package calculator

import (
	"math"
	"testing"
)

// epsilon is the tolerance for floating-point comparisons
const epsilon = 1e-9

// almostEqual safely compares two float64 numbers
func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 2.5, 3.5, 6.0},
		{"negative numbers", -2.5, -3.5, -6.0},
		{"mixed signs", -2.5, 3.5, 1.0},
		{"zeros", 0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); !almostEqual(got, tt.want) {
				t.Errorf("Add(%f, %f) = %f; want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"standard subtraction", 10.0, 4.0, 6.0},
		{"resulting in negative", 4.0, 10.0, -6.0},
		{"subtracting negative", 5.0, -3.0, 8.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.a, tt.b); !almostEqual(got, tt.want) {
				t.Errorf("Subtract(%f, %f) = %f; want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 3.0, 4.0, 12.0},
		{"by zero", 5.0, 0.0, 0.0},
		{"negative times positive", -3.0, 4.0, -12.0},
		{"negative times negative", -3.0, -4.0, 12.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.a, tt.b); !almostEqual(got, tt.want) {
				t.Errorf("Multiply(%f, %f) = %f; want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"even division", 10.0, 2.0, 5.0},
		{"fractional result", 5.0, 2.0, 2.5},
		{"negative division", -10.0, 2.0, -5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.a, tt.b); !almostEqual(got, tt.want) {
				t.Errorf("Divide(%f, %f) = %f; want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		name string
		val  float64
		want float64
	}{
		{"perfect square", 16.0, 4.0},
		{"decimal square", 2.25, 1.5},
		{"zero", 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareRoot(tt.val); !almostEqual(got, tt.want) {
				t.Errorf("SquareRoot(%f) = %f; want %f", tt.val, got, tt.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		val  int
		want int
	}{
		{"base case 0", 0, 1},
		{"base case 1", 1, 1},
		{"standard positive", 5, 120},
		{"larger positive", 7, 5040},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.val); got != tt.want {
				t.Errorf("Factorial(%d) = %d; want %d", tt.val, got, tt.want)
			}
		})
	}
}

func TestNaturalLog(t *testing.T) {
	tests := []struct {
		name string
		val  float64
		want float64
	}{
		{"log of E", math.E, 1.0},
		{"log of 1", 1.0, 0.0},
		{"log of E squared", math.Exp(2), 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NaturalLog(tt.val); !almostEqual(got, tt.want) {
				t.Errorf("NaturalLog(%f) = %f; want %f", tt.val, got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name string
		base float64
		exp  float64
		want float64
	}{
		{"positive exponent", 2.0, 3.0, 8.0},
		{"zero exponent", 5.0, 0.0, 1.0},
		{"negative exponent", 2.0, -1.0, 0.5},
		{"fractional base", 0.5, 2.0, 0.25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Power(tt.base, tt.exp); !almostEqual(got, tt.want) {
				t.Errorf("Power(%f, %f) = %f; want %f", tt.base, tt.exp, got, tt.want)
			}
		})
	}
}
