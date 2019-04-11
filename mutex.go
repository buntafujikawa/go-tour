package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

// Lock と Unlock で囲むことで排他制御で実行するコードを定義できます
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekeysomekey")
	}
	
	time.Sleep(time.Second)
	fmt.Println(c.value("somekey"))
}
