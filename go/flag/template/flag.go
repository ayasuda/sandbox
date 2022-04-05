package main

import (
	"flag"
)

// 構造体には取得用のメソッドを用意すべきか？　面倒だからやらない。別にグローバルにアクセス可能でよかろう。
// その上でコンパイルエラーなどを防ぐために、定義は必須にする。
// つまり動的なメソッドなどは作らない
// 初期値は `default: "hoge"` で取得可能にしようぜ。
type Command struct {
	VerboseLevel uint8  `flag:"verbose"`
	Output       string `flag:"output"`
	Version      bool   `flag:"version"`
}

func NewCommand() *Command {
	return &Command{}
}

// parse は何を返すのか。
// flags, args, subcommands をパースして、最終的に構造体を返せば良い・・・はず？
func Parse(args []string) *Command {
	cmd := NewCommand()
	// TOOD: flag.ExitOnError で良いのか？
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)

	Assgin(fs, cmd)

	fs.Parse(args)
	return cmd
}
func Assgin(fs *flag.FlagSet, cmd *Command) {
	fs.BoolVar(&cmd.Version, "version", false, "dispaly program version")
}

/*
func Assgin(fs *flag.FlagSet, cmd *Command) {
	cv := reflect.ValueOf(cmd)
	t := reflect.TypeOf(*cmd)
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i).Tag.Get("flag")
		if ft != "" {
			fld := cv.Field(i)
			if ft == "version" {
				fp := (*bool)(unsafe.Pointer(fld.Pointer()))
				//fs.BoolVar(&cmd.Version, "version", false, "dispaly program version")

				fs.BoolVar(fp, "version", false, "display program version")
				//fs.BoolVar(&v, "version", false, "display program version")
			}
		}
	}
}
*/
