// Calculator in the terminal
package main

import (
	"fmt"
	"go-calculator/go/src/calc"
	"go-calculator/go/src/keyboard"
	"log"
)

func main() {
	fmt.Println("Enter in your first number")
	number1, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Choose your operator( + - * / ) :")
	operator, err := keyboard.GetString()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter in your second number")
	number2, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	// Start the calculation switch statement
	switch operator {
	case "+":
		sum, err := calc.Add(number1, number2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%0.2f", sum)
		}
	case "-":
		difference, err := calc.Subtract(number1, number2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%0.2f", difference)
		}
	case "*":
		product, err := calc.Multiply(number1, number2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%0.2f", product)
		}
	case "/":
		quotient, err := calc.Divide(number1, number2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%0.2f", quotient)
		}
	default:
		fmt.Println("Invalid Operator")
	}
}
