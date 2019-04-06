package main

import "fmt"

func main() {
	i, j := 42, 2701

	// & オペレータは、そのオペランド( operand )へのポインタを引き出します。
	p := &i
	// * オペレータは、ポインタの指す先の変数を示します。
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
