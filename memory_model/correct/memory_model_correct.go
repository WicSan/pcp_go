package main

var c = make(chan string, 1)

func setup() {
	c <- "hello, world"
}

func main() {
	go setup()
	print(<-c)
}
