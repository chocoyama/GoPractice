package main

import (
    "log"
    "net/http"
    "text/template"
    "path/filepath"
    "sync"
)

type templateHandler struct {
    once sync.Once
    filename string
    templ *template.Template
}

// http.Handlerインターフェースに適合している
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // sync.Once.Doで1回だけの実行を保証することができる
    t.once.Do(func() {
        // ファイルを読み込み、templateをコンパイルして実行
        t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
    })
    t.templ.Execute(w, nil) // templateをResponseWriterに出力
}

func main() {
    // net/httpパッケージを利用し、ルートパスへのリクエストを待ち受ける
    http.Handle("/", &templateHandler{filename: "chat.html"})
    /*
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // リクエストを受け取るとハードコードされたHTMLを返す
        w.Write([]byte(`
            <html>
                <head>
                    <title>チャット</title>
                </head>
                <body>
                    チャットしましょう！
                </body>
            </html>
        `))
    })
    */
    // ListenAndServeメソッドで、8080ポートでWebサーバーを開始する
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
