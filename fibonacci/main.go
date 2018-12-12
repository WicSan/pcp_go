//Endrekursive Fibonacci-Berechnungmit Ein-und Ausgabe
//2.Aufgabe Scheme 5+6
package main

import (
	"fmt"
)

func fibonacci() func() int {
	f_1 := 0
	f_2 := 1
	return func() int {
		f_2, f_1 = f_1, f_2+f_1
		return f_2
	}
}

func main() {
	var input int
	fmt.Print("Gib eine Zahl ein: ")
	fib := fibonacci()
	_, err := fmt.Scanf("%d", &input)

	if err == nil {
		for i := 0; i < input; i++ {
			fib()
		}
		fmt.Printf("Die %d. Fibonacci-Zahl ist %d", input, fib())
	}
}
