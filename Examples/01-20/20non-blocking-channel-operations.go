package main

import "fmt"

// 在 Go 里，select 的 default case 用于 非阻塞选择，
// 也就是当没有任何 channel 准备好时，
// select 会立即执行 default，而不是阻塞等待。
func main() {
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
