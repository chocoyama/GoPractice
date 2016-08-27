package main

import (
    "fmt"
)

/*
スライスはlengthとcapacityを持つ
length: 配列が持っている要素の数
capacity: 元となる配列の要素の数

:nは元となる配列のn番目の要素まで（n番目は含まない）のスライスを作る
    capacityは変更されない
:nは元となる配列のn番目の要素から最後の要素まで（n番目を含む）のスライスを作る
    除外された先頭要素分capacityも変更される
*/
func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
    s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    printSlice(s) // => len=10 cap=10 [1 2 3 4 5 6 7 8 9 10]

    s = s[:0]
    printSlice(s) // => len=0 cap=10 []
    s = s[:5]
    printSlice(s) // => len=5 cap=10 [1 2 3 4 5]

    s = s[1:]
    printSlice(s) // => len=4 cap=9 [2 3 4 5]
    s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    s = s[1:]
    printSlice(s) // => len=9 cap=9 [2 3 4 5 6 7 8 9 10]
    s = s[:4]
    printSlice(s) // => len=4 cap=9 [2 3 4 5]
    s= s[1:]
    printSlice(s) // => len=3 cap=8 [3 4 5]

    s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    s = s[2:5]
    printSlice(s) // => len=3 cap=8 [3 4 5]

    /*
    nil slice
    */
    var nil_s []int
    printSlice(nil_s) // => len=0 cap=0 []
}
