package main

import (
	"fmt"
	"scientific-calculator/calculator"
)

func main() {
	// commit
	// test commit
	// test commit 
	var choice int

	for {
		fmt.Println("\n--- Scientific Calculator ---")
		fmt.Println("1. Addition (+)")
		fmt.Println("2. Subtraction (-)")
		fmt.Println("3. Multiplication (*)")
		fmt.Println("4. Division (/)")
		fmt.Println("5. Square Root (√x)")
		fmt.Println("6. Factorial (x!)")
		fmt.Println("7. Natural Logarithm (ln(x))")
		fmt.Println("8. Power Function (x^b)")
		fmt.Println("9. Exit")
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if choice == 9 {
			fmt.Println("Exiting calculator. Goodbye!")
			break
		}

		processChoice(choice)
	}
}

func processChoice(choice int) {
	var a, b float64

	switch choice {
	case 1:
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)
		fmt.Print("Enter second number: ")
		fmt.Scan(&b)
		result := calculator.Add(a, b)
		fmt.Printf("Result: %f\n", result)

	case 2:
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)
		fmt.Print("Enter second number: ")
		fmt.Scan(&b)
		result := calculator.Subtract(a, b)
		fmt.Printf("Result: %f\n", result)

	case 3:
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)
		fmt.Print("Enter second number: ")
		fmt.Scan(&b)
		result := calculator.Multiply(a, b)
		fmt.Printf("Result: %f\n", result)

	case 4:
		fmt.Print("Enter numerator: ")
		fmt.Scan(&a)
		fmt.Print("Enter denominator: ")
		fmt.Scan(&b)
		if b == 0 {
			fmt.Println("Error: Cannot divide by zero.")
		} else {
			result := calculator.Divide(a, b)
			fmt.Printf("Result: %f\n", result)
		}

	case 5:
		fmt.Print("Enter value for x: ")
		fmt.Scan(&a)
		if a < 0 {
			fmt.Println("Error: Cannot calculate the square root of a negative number.")
		} else {
			result := calculator.SquareRoot(a)
			fmt.Printf("Result: %f\n", result)
		}

	case 6:
		var n int
		fmt.Print("Enter integer value for x: ")
		fmt.Scan(&n)
		result := calculator.Factorial(n)
		if result == -1 {
			fmt.Println("Error: Factorial is not defined for negative numbers.")
		} else {
			fmt.Printf("Result: %d\n", result)
		}

	case 7:
		fmt.Print("Enter value for x: ")
		fmt.Scan(&a)
		if a <= 0 {
			fmt.Println("Error: Natural logarithm is only defined for positive numbers.")
		} else {
			result := calculator.NaturalLog(a)
			fmt.Printf("Result: %f\n", result)
		}

	case 8:
		fmt.Print("Enter base (x): ")
		fmt.Scan(&a)
		fmt.Print("Enter exponent (b): ")
		fmt.Scan(&b)
		result := calculator.Power(a, b)
		fmt.Printf("Result: %f\n", result)

	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 9.")
	}
}
