package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map はキーと値とを関連付けます
var m map[string]Vertex

// mapリテラルは、structリテラルに似ていますが、 キー ( key )が必要です
var n = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(n)

	// mutating maps
	o := make(map[string]int)

	o["Answer"] = 42
	fmt.Println("The value:", o["Answer"])

	o["Answer"] = 48
	fmt.Println("The value:", o["Answer"])

	delete(o, "Answer")
	fmt.Println("The value:", o["Answer"])

	// キーに対する要素が存在するかどうかは、2つの目の値で確認します
	v, ok := o["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
