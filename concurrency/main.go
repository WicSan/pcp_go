package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	channel := make(chan int, 2)
	channel2 := make(chan string, 1)

	fmt.Println("Now waiting for things to happen!")
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Print(3000)
		channel <- 3000
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Print(6000)
		channel <- 6000
	}()

	go func() {
		var sum int
		for i := 0; i < 2; i++ {
			sum += <-channel
		}
		time.Sleep(2 * time.Second)
		channel2 <- ("was waiting for " + strconv.Itoa(sum+2000) + "ms")
	}()

	for {
		select {
		case msg := <-channel2:
			fmt.Println(msg)
			fmt.Println("Done.")
			return
		default:
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
