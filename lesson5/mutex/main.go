package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	v map[string]int
}

func (c *Counter) Inc(key string) {
	c.Lock()
	c.v[key]++
	c.Unlock()
}

func (c *Counter) Value(key string) int {
	defer c.Unlock()
	c.Lock()
	return c.v[key]
}

func main() {
	c := Counter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.Inc("key")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("key"))
}
