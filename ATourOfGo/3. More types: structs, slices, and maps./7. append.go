package main

import "fmt"

/*
元の配列が変数群を追加する際に容量が小さい場合は、
現在の容量を2倍にしたサイズの配列を割りあて直す
*/
func main() {
    var s []int
    printSlice(s) // => len=0 cap=0 []

    s = append(s, 0)
    printSlice(s) // => len=1 cap=1 [0]

    s = append(s, 1)
    printSlice(s) // => len=2 cap=2 [0 1]

    s = append(s, 2)
    printSlice(s) // => len=3 cap=4 [0 1 2] 足りなくなったので2倍の容量確保された！
    s = append(s, 3)
    printSlice(s) // => len=4 cap=4 [0 1 2 3]
    s = append(s, 4)
    printSlice(s) // => len=5 cap=8 [0 1 2 3 4] 足りなくなったので2倍の容量確保された！

    s = append(s, 5, 6, 7)
    printSlice(s) // => len=8 cap=8 [0 1 2 3 4 5 6 7 8]
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
