package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {
	op := flag.String("op", "add", "Operation type (add, subtract, multiply, divide, exponent, modulus)")
	a := flag.Float64("a", 0, "Operand 1")
	b := flag.Float64("b", 0, "Operand 2")
	flag.Parse()

	if *op == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	result, err := calculate(*op, *a, *b)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %.4f\n", result)

}

func calculate(operation string, a, b float64) (float64, error) {
	switch operation {
	case "add":
		return addition(a, b), nil
	case "subtract":
		return subtraction(a, b), nil
	case "multiply":
		return multiplication(a, b), nil
	case "divide":
		return division(a, b)
	case "exponent":
		return exponentate(a, b), nil
	case "modulus":
		return modulus(a, b)
	default:
		return 0, fmt.Errorf("Operation %s not reckognized. Please try again.\n", operation)
	}
}

func addition(a, b float64) float64 {
	return a + b
}

func subtraction(a, b float64) float64 {
	return a - b
}

func multiplication(a, b float64) float64 {
	return a * b
}

func division(a, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return 0, errors.New("Division by zero is not allowed.")
	}
}

func exponentate(a, b float64) float64 {
	return math.Pow(a, b)
}

func modulus(a, b float64) (float64, error) {
	if b != 0 {
		return math.Mod(a, b), nil
	} else {
		return 0, errors.New("Modulus by zero is not allowed.")
	}
}
