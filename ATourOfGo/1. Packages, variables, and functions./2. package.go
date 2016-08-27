package main // プログラムはmainパッケージから始まる

/*
パッケージ名はインポートパスの最後の要素と同じ名前になる。
math/randの場合、package randステートメントで始まるファイル群で構成する
※ URLを含むインポートパスが"golang.org/x/net/websocket"だった場合は、package websocketになる
*/
import(
    "fmt"
    "math/rand"
)

func main() {
    /*
    パッケージをインポートすると、そのパッケージがエクスポートしている名前を参照することができる。
    最初の名前が大文字で始まる名前は、外部のパッケージから参照できる公開された名前（exported name）。
    */
    fmt.Println("My favorite number is", rand.Intn(10))
}
