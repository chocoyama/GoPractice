package main

import (
    "fmt"
)

/*
定数はconstキーワードを使って変数同じように宣言する（関数内外を問わない）
文字、文字列、真偽値、数値のみで使える
「:=」を使った宣言はできない
*/
const Pi = 3.14

const (
    Big = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int {
    return x * 10 + 1
}

func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    const World = "世界"
    fmt.Println("Hello", World) // => Hello 世界
    fmt.Println("Happy", Pi, "Day") // => Happy 3.14 Day

    const Truth = true
    fmt.Println("Go rules?", Truth) // => Go rules? true

    fmt.Println(needInt(Small)) // => 21
    fmt.Println(needFloat(Small)) // => 0.2
    fmt.Println(needFloat(Big)) // => 1.2676506002282295e+29
}
