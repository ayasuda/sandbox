package main

import (
	"strings"
	"testing"
)

func TestPrse(t *testing.T) {
	tcs := []struct {
		name      string
		input     string
		assertion func(t *testing.T, cmd *Command)
	}{
		{
			name:  "no argument: --hoge",
			input: "",
			assertion: func(t *testing.T, cmd *Command) {
				if cmd.Version == true {
					t.Errorf("expect Version set false, but got true")
				}
			},
		},
		{
			name:  "show version",
			input: "--version",
			assertion: func(t *testing.T, cmd *Command) {
				if cmd.Version != true {
					t.Errorf("expect Version set true, but got false")
				}
			},
		},
		//		{
		//			name:  "normal",
		//			input: "--hoge fuga piyo piyo",
		//			assertion: func(t *testing.T, cmd *Command) {
		//			},
		//		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			cmd := Parse(strings.Split(tc.input, " "))
			t.Logf("argument: '%s'", tc.input)
			tc.assertion(t, cmd)
		})
	}
	t.Skip()
}

// サブコマンドがなければサブコマンド用のヘルプ
// cmd --version はバージョン表記
// cmd sub args...
// cmd -v sub args...
// cmd --opt sub --opt args...
// を、どう取り扱うのか。
// 結果を構造体として取り扱いたい
// in := parse()
// in.SubCommand()
// in.Argument()
// in.Option("v")
// in.SubCommand().Option()

//"--opt-for-parent sub --opt-for-sub subArgs..."
