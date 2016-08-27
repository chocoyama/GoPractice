package main

import (
    "fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    /*
    i: インデックス
    v: インデックスの場所の要素のコピー
    */
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    /*
    インデックスや値は、"_"へ代入することで捨てられる
    */
    pow2 := make([]int, 10)
    for i := range pow2 {
        pow2[i] = 1 << uint(i)
    }
    for _, value := range pow2 {
        fmt.Printf("%d\n", value)
    }
}
