package main

import (
	"fmt"
)

func PrintNumbersAndSquares(liste []int) {
	for _, value := range liste {
		var x = value * value
		fmt.Println(value, x)
	}
}

func main() {
	numbers := []int{1, 2, 3, 5, 8}
	PrintNumbersAndSquares(numbers)
}