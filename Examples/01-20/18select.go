package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println(time.Now(), "received", msg1)
		case msg2 := <-c2:
			fmt.Println(time.Now(), "received", msg2)
		}
	}
}
