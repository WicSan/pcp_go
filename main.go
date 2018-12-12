package main

import "fmt"

func fibonacci() func() int {
	f_1 := 0
	f_2 := 1
	return func() int {
		f_2, f_1 = f_1, f_2+f_1
		return f_2
	}
}

func main() {
	fibN := readInt()
	i := 0

	f := fib()

	for i < fibN {
		fmt.Print(f(), " ")
		i++
	}
}
