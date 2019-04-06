package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main()  {
    // for ループに利用する range は、スライスや、マップ( map )をひとつずつ反復処理するために使います。
    // rangeは反復毎に2つの変数を返します
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
    
    pow2 := make([]int, 10)
    for i := range pow2 {
        pow2[i] = 1 << uint(i) // == 2**i
    }
    
    //  " _ "(アンダーバー) へ代入することで捨てることができます
    // もしインデックスだけが必要なのであれば、 " , value " を省略
    for _, value := range pow2 {
        fmt.Printf("%d\n", value)
    }
}
