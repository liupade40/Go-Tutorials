package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
	}

	// 启动3个并发任务
	wg.Add(3)

	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	go func() {
		defer wg.Done()
		doIncrement("b", 10000)
	}()

	wg.Wait()
	fmt.Println(c.counters)
}
