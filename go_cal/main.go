package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	n := time.Now()
	f := time.Date(n.Year(), n.Month(), 1, 0, 0, 0, 0, time.Local)
	l := lastDay(n)

	ms := n.Month().String()
	ls := (15 - len(ms)) / 2
	fmt.Printf("%s%s %d\n", strings.Repeat(" ", ls), ms, n.Year())

	var idt, cnt int
	switch f.Weekday() {
	case time.Sunday:
		idt = 0
		cnt = 6
	case time.Monday:
		idt = 3
		cnt = 5
	case time.Tuesday:
		idt = 6
		cnt = 4
	case time.Wednesday:
		idt = 9
		cnt = 3
	case time.Thursday:
		idt = 12
		cnt = 2
	case time.Friday:
		idt = 15
		cnt = 1
	case time.Saturday:
		idt = 18
		cnt = 0
	}
	fmt.Println("Su Mo Tu We Th Fr Sa")
	fmt.Printf("%s", strings.Repeat(" ", idt))
	for i := 1; i <= l.Day(); i++ {
		if i < 10 {
			fmt.Printf(" %d", i)
		} else {
			fmt.Printf("%d", i)
		}
		if cnt <= 0 {
			fmt.Printf("\n")
			cnt = 6
		} else {
			fmt.Printf(" ")
			cnt--
		}
	}
	fmt.Printf("\n")
}

func lastDay(n time.Time) time.Time {
	var f time.Time
	if n.Month() == time.November {
		f = time.Date(n.Year()+1, time.January, 1, 0, 0, 0, 0, time.Local)
	} else {
		f = time.Date(n.Year(), n.Month()+1, 1, 0, 0, 0, 0, time.Local)
	}
	return f.Add(time.Hour * 24 * -1)
}
