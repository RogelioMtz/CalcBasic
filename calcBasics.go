package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Constants for operations
const (
	Add         = "+"
	Subtract    = "-"
	Multiply    = "*"
	Divide      = "/"
	Exponential = "^"
	SquareRoot  = "v"
	Quit        = "q"
)

// Main function
func main() {
	for {
		operation := getOperation()
		if operation == Quit {
			fmt.Println("Quitting the program.")
			break
		}
		x, y := getUserInput(operation)
		result, err := performOperation(x, y, operation)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}

// Function to get the operation from the user
func getOperation() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`
				Choose operation
                 Basic Operations:       Advanced Operations:
                 (+) Add                 (^) Exponential
                 (-) Subtract            (v) Square Root
                 (*) Multiply
                 (/) Divide              (q) Quit`)
	operation, _ := reader.ReadString('\n')
	return strings.TrimSpace(operation)
}

// Function to get user input based on the operation
func getUserInput(operation string) (float64, float64) {
	reader := bufio.NewReader(os.Stdin)
	var x, y float64

	fmt.Println("Enter first number: ")
	x = readFloat(reader)

	if operation != SquareRoot {
		fmt.Println("Enter second number: ")
		y = readFloat(reader)
	}

	return x, y
}

// Function to read a float64 from user input
func readFloat(reader *bufio.Reader) float64 {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number: ")
		} else {
			return value
		}
	}
}

// Function to perform the selected operation
func performOperation(x, y float64, operation string) (float64, error) {
	switch operation {
	case Add:
		return add(x, y), nil
	case Subtract:
		return subtract(x, y), nil
	case Multiply:
		return multiply(x, y), nil
	case Divide:
		return divide(x, y)
	case Exponential:
		return exponential(x, y), nil
	case SquareRoot:
		result, err := squareRoot(x)
		if err != nil {
			return 0, err
		}
		return result, nil
	case Quit:
		fmt.Println("Quitting the program.")
		os.Exit(0)
		return 0, nil // This line will never be reached, but it satisfies the return requirement
	default:
		return 0, fmt.Errorf("invalid operation")
	}
}

// Functions for basic operations
func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func exponential(a, b float64) float64 {
	return math.Pow(a, b)
}

func squareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("cannot calculate square root of a negative number")
	}
	return math.Sqrt(a), nil
}
