package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

// メソッドが変数レシーバである場合、呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができます
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 変数の引数を取る関数は、特定の型の変数を取る必要があります
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// pointer receivers
	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v2.Abs())

	v3 := Vertex{3, 4}
	Scale(&v3, 10)
	fmt.Println(Abs(v3))

	// Methods and pointer indirection 
	v4 := Vertex{3, 4}
	v4.Scale(2) // v.Scale(2) のステートメントを (&v).Scale(2) として解釈
	Scale(&v4, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	Scale(p, 8)

	fmt.Println(v4, p)

	v5 := Vertex{3, 4}
	fmt.Println(v5.Abs())
	fmt.Println(Abs(v5))

	p2 := &Vertex{4, 3}
	fmt.Println(p2.Abs()) // p.Abs() は (*p).Abs() として解釈されます
	fmt.Println(Abs(*p2))
}
