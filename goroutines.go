package main

import (
	"fmt"
	"time"
)

func main() {
	go say("world")
	say("hello")

	// channels
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// buffered channel
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)

	// range and close
	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}
	
	// select
	c3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c3)
		}
		quit <- 0
	}()
	fibonacci2(c3, quit)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// channels
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// チャネル( Channel )型は、チャネルオペレータの <- を用いて値の送受信ができる通り道です
	c <- sum
}

// range and close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// select
// select は、複数ある case のいずれかが準備できるようになるまでブロックし、準備ができた case を実行します
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		//default:
		//	// どのcaseも準備できない場合
		//	fmt.Println("default")
		}
	}
}
