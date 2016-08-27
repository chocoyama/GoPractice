package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

/*
"func"と関数名の間にstructを指定することで、メソッドを追加できる
 ※ このとき、レシーバの元の変数のコピーを操作して実行される
*/
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

/*
通常の関数として記述した場合
*/
func Abs(v Vertex) float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

/*
typealiasの記述をしておけば、任意の型でメソッド宣言することができる
*/
type MyFloat float64
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    } else {
        return float64(f)
    }
}

/*
レシーバのフィールド値を変更するようなメソッドを作りたいときは、ポインタで指定する
もしくはレシーバと同様の型の新しいインスタンスを返すでも良さそう
*/
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

/*
メソッド呼び出す時のレシーバは変数、ポインタどちらでもよく、Goが勝手に解釈してくれる
    v.Abs() => OK
    v.Scale() => OK
    p := &v
    p.Abs() => OK
    p.Scale() => OK
*/
func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
    fmt.Println(Abs(v))

    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())

    v.Scale(10)
    fmt.Println(v.Abs())
}
