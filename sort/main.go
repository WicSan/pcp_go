//6.Aufgabe Scheme 5+6
//diffrence slices([]) & arrays([3] fixed size)
package main

import "fmt"

func calcList(functions []func(int, int) int, a, b int) {
	for _, f := range functions {
		fmt.Println(f(a, b))
	}

	fmt.Println("finished")
}

func main() {
	var functions [2]func(int, int) int
	functions[0] = func(a, b int) int {
		return a * b
	}

	functions[1] = func(a, b int) int {
		return (a + b) * 2
	}

	calcList(functions[:], 2, 3)
}
