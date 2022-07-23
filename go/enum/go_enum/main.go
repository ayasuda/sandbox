package main

import (
	"fmt"
	"time"
)

type Gengo int

const (
	MEIJI Gengo = iota
	TAISHO
	SHOWA
	HEISEI
	REIWA
	UNKWON
)

var gengoDict = map[Gengo]struct {
	name     string
	alphabet string
	start    time.Time
}{
	MEIJI:  {name: "明治", alphabet: "M", start: time.Date(1868, 1, 25, 0, 0, 0, 0, time.Local)},
	TAISHO: {name: "大正", alphabet: "T", start: time.Date(1912, 7, 30, 0, 0, 0, 0, time.Local)},
	SHOWA:  {name: "昭和", alphabet: "S", start: time.Date(1926, 12, 25, 0, 0, 0, 0, time.Local)},
	HEISEI: {name: "平成", alphabet: "H", start: time.Date(1989, 1, 8, 0, 0, 0, 0, time.Local)},
	REIWA:  {name: "令和", alphabet: "R", start: time.Date(2019, 5, 1, 0, 0, 0, 0, time.Local)},
	UNKWON: {name: "-", alphabet: "-", start: time.Date(9999, 12, 31, 0, 0, 0, 0, time.Local)},
}

func (g Gengo) String() string {
	return gengoDict[g].name
}

func (g Gengo) Alphabet() string {
	return gengoDict[g].alphabet
}

func (g Gengo) Start() time.Time {
	return gengoDict[g].start
}

func From(t time.Time) Gengo {
	for _, g := range []Gengo{MEIJI, TAISHO, SHOWA, HEISEI, REIWA} {
		if t.Before(g.Start()) {
			return g
		}
	}
	return UNKWON
}

func main() {
	a := From(time.Date(2001, 1, 1, 0, 0, 0, 0, time.Local))
	fmt.Printf("Gengo name is %s, and alphabet is %s\n", a.String(), a.Alphabet())
}
