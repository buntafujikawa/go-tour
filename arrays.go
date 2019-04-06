package main

import "fmt"

// [n]T 型は、型 T の n 個の変数の配列( array )を表します
func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
    
    slices()
    references()
    sliceLiterals()
}

// これは最初の要素を含む half-open レンジを選択しますが、最後の要素は省かれます。
// 次の式は a の要素の内 1 から 3 を含むスライスを作ります。
func slices() {
    primes := [6]int{2, 3, 5, 7, 11, 13}
    var s []int = primes[1:4]
    fmt.Println(s)
}

func references() {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }

    fmt.Println(names)
    
    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)

    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)
}

func sliceLiterals() {
    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(q)

    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)

    s := []struct {
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {5, true},
        {7, true},
        {11, false},
        {13, true},
    }
    fmt.Println(s)
}
