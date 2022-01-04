// Package calc has add, subtract, multiply, divide between math calculations between two numbers.
package calc

import "errors"

// Add finds the sum of two numbers
// Returns a sum and any error encountered.
func Add(summand1 float64, summand2 float64) (float64, error) {
	return summand1 + summand2, nil
}

// Subtract finds the difference of two numbers.
// Returns a difference and any error encountered.
func Subtract(minuend float64, subtrahend float64) (float64, error) {
	return minuend - subtrahend, nil
}

// Multiply finds the product of two numbers.
// Returns the product and any error encountered.
func Multiply(multiplier float64, multiplicand float64) (float64, error) {
	return multiplier * multiplicand, nil
}

// Divide finds the division of two numbers.
// Returns the products and any error encountered.
func Divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}
