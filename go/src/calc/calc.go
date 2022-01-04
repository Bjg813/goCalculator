package calc

import "errors"

func Add(summand1 float64, summand2 float64) (float64, error) {
	return summand1 + summand2, nil
}

func Subtract(minuend float64, subtrahend float64) (float64, error) {
	return minuend - subtrahend, nil
}

func Multiply(multiplier float64, multiplicand float64) (float64, error) {
	return multiplier * multiplicand, nil
}

func Divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}
