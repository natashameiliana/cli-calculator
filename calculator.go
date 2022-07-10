package main

import (
	"flag"
	"fmt"
)

func sum(first int, second int) (result int) {
	result = first + second
	return
}

func substract(first int, second int) (result int) {
	result = first - second
	return
}

func multiply(first int, second int) (result int) {
	result = first * second
	return
}

func divided(first int, second int) (result int) {
	result = first / second
	return
}

func modulo(first int, second int) (result int) {
	result = first % second
	return
}

func main() {
	// Calculator
	var first int
	var op string
	var second int

	flag.IntVar(&first, "first", 0, "-first= ")
	flag.StringVar(&op, "op", "", "-op= ")
	flag.IntVar(&second, "second", 0, "-second= ")

	flag.Parse()

	switch {
	case op == "+":
		total := sum(first, second)
		fmt.Printf("Result of %d + %d = %d \n", first, second, total)
	case op == "-":
		total := substract(first, second)
		fmt.Printf("Result of %d - %d = %d \n", first, second, total)
	case op == "*":
		total := multiply(first, second)
		fmt.Printf("Result of %d * %d = %d \n", first, second, total)
	case op == "/":
		total := divided(first, second)
		fmt.Printf("Result of %d / %d = %d \n", first, second, total)
	case op == "%":
		total := modulo(first, second)
		fmt.Printf("Result of %d % %d = %d \n", first, second, total)
	default:
		fmt.Println("Silahkan coba lagi \n")
	}

}

