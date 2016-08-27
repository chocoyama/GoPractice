package main

import(
    "fmt"
    "math/cmplx"
)

/*
Golangの基本型は次のとおり
- bool
- string
- int int8 int16 int32(rune) int64
- uint uint8(byte) uint16 uint32 uint64 uintptr
- float32 float64
- complex64 complex128

int uint uintptrは、32bitシステムでは32bit, 64bitシステムでは64bit
*/

var (
    ToBe bool = false
    MaxInt uint64 = 1<<64 -1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

/*
変数に初期値を与えずに宣言すると、ゼロ値（zero value）が与えられる
- 数値型: 0
- bool型: flase
- string型: ""（空文字）
*/
var (
    i int // 0
    f64 float64 // 0
    b bool // false
    s string // ""
)

func main() {
    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe) // => bool(false)
    fmt.Printf(f, MaxInt, MaxInt) // => uint64(18446744073709551615)
    fmt.Printf(f, z, z) // => complex128((2+3i))

    fmt.Printf("%v %v %v %q\n", i, f64, b, s) // => 0 0 false ""
}
