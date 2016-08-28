## $GOPATH
GOPATHには2つの役割がある
1. ビルド時のインポートパスとして、GOPATHに指定した全てのディレクトリを参照する
2. go getコマンドで外部パッケージを読み込んだ時、GOPATHの先頭のディレクトリにダウンロードする

## GOPATH指定のベストプラクティス
自身のプロジェクトを書く場合もGOPATHの元で書くと良い
ex) $GOPATH/src/github.com/foo/bar

ただし、ディレクトリを1つだけ指定してしまうと、サードパーティのライブラリと自分が書いたソースコードが混ざってしまう。
これを解決するには、GOPATHの性質を利用して、先頭のディレクトリにはサードパーティーのパスを指定し、続けて自身の作業ディレクトリを設定ておくのがよい
```
$ export GOPATH=$HOME/go/third-party:$HOME/go/my-project
$ export PATH=$HOME/go/third-party/bin:$HOME/go/my-project/bin:
```
