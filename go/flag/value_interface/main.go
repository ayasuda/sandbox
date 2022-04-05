package main

import (
	"flag"
	"fmt"
)

type Foo struct {
	Attr string
}

// Value interface を満たすために String() を実装する
// String() 関数はプログラム診断などのために使います
func (f *Foo) String() string {
	return f.Attr
}

// Value interface を満たすために Set() を実装する
// Set() 関数は flag.Parse が実際のコマンドライン引数をこの変数に割り当てるために使われます
func (f *Foo) Set(s string) error {
	f.Attr = s
	return nil
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuseday
	Wednesday
	Thursday
	Friday
	Saturday
	UNKOWN
)

func (w *Weekday) String() string {
	return ""
}

func (w *Weekday) Set(s string) error {
	w = &Monday
	return nil
}

func main() {
	var foo Foo
	flag.Var(&foo, "foo", "help message")
	flag.Parse()
	fmt.Println("foo is", foo.String())
}
