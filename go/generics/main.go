package main

import "fmt"

func Reverse[T any](s []T) {
	f := 0
	l := len(s) - 1
	for f < l {
		s[f], s[l] = s[l], s[f]
		f++
		l--
	}
}

type User struct {
	Name string
}

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("ordered: ", ints)
	Reverse(ints)
	fmt.Println("reversed: ", ints)

	strings := []string{"apple", "grape", "orange", "banana"}

	fmt.Println("ordered: ", strings)
	Reverse(strings)
	fmt.Println("reversed: ", strings)

	users := []User{{"佐藤"}, {"鈴木"}, {"高橋"}, {"田中"}, {"伊藤"}}

	fmt.Println("ordered: ", users)
	Reverse(users)
	fmt.Println("reversed: ", users)
}
