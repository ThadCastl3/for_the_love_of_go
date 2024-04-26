// package calculator does simple calculations
package calculator

import (
	"errors"
	"math"
)

// add takes two numbers and returns the result of adding
// them together
func Add(a, b float64) float64 {
	return a + b
}

// subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// multiply takes two numbers a and b, and
// returns the result of multiplying them.
func Multiply(a, b float64) float64 {
	return a * b
}

// divide takes two numbers a and b, and
// checks if b is not zero, then returns the result of dividing a by b.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// sqrt takes a number a and returns the square root of a.
// It returns an error if a is negative
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot square root negative number")
	}
	return math.Sqrt(a), nil
}
