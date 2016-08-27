package main

import (
    "fmt"
    "math"
)

/*
T(v)は変数vをT型へ変換する
*/
func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x * x + y * y))
    var z uint = uint(f)

    i := 42 // int
    f := 3.142 // float64
    g := 0.867 + 0.5i // complex128

    fmt.Println(x, y, z) // => 3 4 5
}
