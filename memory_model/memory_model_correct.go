package main

var a string
var c = make(chan int, 10)

func f() {
    <-c
    print(a)
}

func hello() {
	a = "hello, world"
	go f()
    c <- 0
}

func main() {
	hello()
}
