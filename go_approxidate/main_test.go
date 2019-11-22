package main

import (
	"testing"
)

func Test_apprxidate(t *testing.T) {
	tt := []struct {
		source string
		expect string
	}{{
		source: "2019-08-23 17:22:27",
		expect: "2019-08-23 17:22:27.000000 +0900 JST m=+0.000000000",
	}}

	for i, tc := range tt {
		p, err := approxidate(tc.source)
		if err != nil {
			t.Errorf("%d, want no err, but got %v", i, err)
		}
		if p.String() != tc.expect {
			t.Errorf("%d,\nwant\t%s,\ngot\t %s", i, tc.expect, p.String())
		}
	}
}
