package main

import (
    "fmt"
    "io"
    "strings"
)

/*
ioパッケージは、データストリームを読むことを表現するio.Readerインターフェースを規定している
func (T) Read(b []byte) (n int, err error)
    データをバイトスライスに格納し、入れたバイトのサイズとエラーの値を返す
    ストリームの終端はio.EOFのエラーで返す


Go標準ライブラリには以下のような実装がある
- interface
- file
- network connection
- compression
- encryption
*/

func main() {
    r := strings.NewReader("Hello, Reader!")

    b := make([]byte, 8)
    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        if err == io.EOF {
            break
        }
    }
}
