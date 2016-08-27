package main

import (
    "fmt"
)

/*
要素の挿入、更新
m[key] = elem

要素の取得
elem = m[key]

要素の削除
delete(m, key)

キーに対する要素の存在チェック
elem, ok := m[key]
    elem: 存在すればその要素、存在しなければmapの要素の型のゼロ値
    ok: bool 存在するかどうか
*/

func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])

    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])

    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
