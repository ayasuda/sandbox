package main

import (
	"fmt"
	"os"
)

// main は何をするのか？
// 通常、 ROOT/cmd/cmdName/main.go に書かれる。
// それ以外のパッケージは ROOT/ におく。なので、 main.go は package DOMAIN/ROOT を読み込むことになる
// https://github.com/golang-standards/project-layout/issues/117 の通り、 golang-standars/project-layout には従わない
// run application
//  read configuration file, environment value and parse flag or optionns.
// あくまでも俺は、フラグの取り扱い方のみのテンプレートを作っているはずだ。
// だから、フラグをパースして、構造体にアサインできれば良いはずだ。
// 構造体からのゲッターなどは今回は考える必要はない
// パースして・・・
func main() {

	Parse(os.Args)
	fmt.Println("hello")
}
