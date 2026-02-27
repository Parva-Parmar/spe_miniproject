package calculator

import (
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	return a / b
}

func SquareRoot(x float64) float64 {
	return math.Sqrt(x)
}

func Factorial(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func NaturalLog(x float64) float64 {
	return math.Log(x)
}

func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}
