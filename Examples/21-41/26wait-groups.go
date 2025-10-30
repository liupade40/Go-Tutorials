package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // 每个 goroutine 对应一次 Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	wg.Wait() // 等所有 goroutine 执行完
}
