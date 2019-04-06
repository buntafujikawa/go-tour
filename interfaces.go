package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
} 

func (t T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i , i)
}

func main() {
	var i I = T{"hello"}
	i.M()
	
	// interface values
	var i2 I
	
	i2 = &T{"Hello"}
	describe(i2)
	i2.M()
	
	i2 = F(math.Pi)
	describe(i2)
	i2.M()
}
