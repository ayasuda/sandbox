package main

import (
	"flag"
	"fmt"
)

func main() {
	// オプションを宣言する
	// 名前は flagname で、数値。例えば、 --flagname 2525 のようになる
	// デフォルト値は 1234 とする
	// オプションの使用方法として "help message for flagname" を記述しておく
	var input = flag.Int("flagname", 1234, "help message for flagname")

	// 与えられた引数をパースする
	// ここでは `ip` に `--flagname` で与えられた数値をセットする
	flag.Parse()
	fmt.Println("input has value ", *input)
}
