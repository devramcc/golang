package main

import (
	"fmt"
)

func main() {
	printMe(10 + 1)
	fmt.Println(intDivision(10, 2))
	fmt.Println(intDivision(10, 3))
	var result, remainder = intDivisionwithremainder(10, 0)
	fmt.Printf("the result of integer division is %v and the remainder is %d", result, remainder)
}

func printMe(treshHold int8) {
	fmt.Println(treshHold)
}

func intDivision(numerator int, denominator int) int {
	return numerator / denominator
}

func intDivisionwithremainder(numerator int, denominator int) (int, int) {
	var err error
	if denominator == 0 {
		err = fmt.Errorf("cannot divide by zer1o")
	}
	if err != nil {
		fmt.Println(err)
	}

	return numerator / denominator, numerator % denominator
}

// Example of switch vs switch var
func switchStatement(i int) {
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}
}

func switchStatementVar(a string, b string, c string) {
	switch {
	case a == "one" && b == "two" && c == "three":
		fmt.Println("one two three")
	case a == "one" && b == "two":
		fmt.Println("one two")
	case a == "one" && c == "three":
		fmt.Println("one three")
	case b == "two" && c == "three":
		fmt.Println("two three")
	default:
		fmt.Println("default")
	}
}
