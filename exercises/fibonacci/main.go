//Endrekursive Fibonacci-Berechnungmit Ein-und Ausgabe
//2.Aufgabe Scheme 5+6
package main

import (
	"fmt"
)

func fibonacci() func() int64 {
	f1 := int64(0)
	f2 := int64(1)
	return func() int64 {
		f2, f1 = f1, f2+f1
		return f2
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
