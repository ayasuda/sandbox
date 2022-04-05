package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Int("sample", 0, "default usage message")

	// Init を呼び出すことで errorHandling が変更できる
	// しかし、CommandLine はデフォルトで ExitOnError が指定されているのでコメントアウトする
	// flag.CommandLine.Init(os.Arg[0], flag.ExitOnError)

	flag.Parse()
	fmt.Println("This line should not be printed if -h, --help added or some error happened")
}
