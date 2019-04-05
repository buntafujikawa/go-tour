package syntax

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 1} // xのみ1で初期化、yは0
    v3 = Vertex{}
    p = &Vertex{1, 2}
)

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
    
    z := Vertex{1, 2}
    p := &z
    p.X = 1e9
    fmt.Println(z)
    
    fmt.Println(v1, p, v2, v3)
}
