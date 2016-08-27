package main

import (
    "fmt"
    "math"
)

/*
インターフェースを実装することを明示的に宣言する必要はない
*/
type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt(2))
    v := Vertex{3, 4}

    a = f
    a = &v

    // a = v => compile error!!（Absメソッドが*Vertexの定義になっているため

    fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    } else {
        return float64(f)
    }
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
