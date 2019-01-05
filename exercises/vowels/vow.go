package main

import (
	"fmt"
	"unicode"
)

var vows = map[rune]struct{}{
	'a':struct{}{},
	'e':struct{}{},
	'i':struct{}{},
	'o':struct{}{},
	'u':struct{}{},
}

func vow1(str string) int {
	n := 0
	for _, r := range str {
		if _, ok := vows[unicode.ToLower(r)]; ok {
			n ++
		}
	}
	return n
}

func main() {
	fmt.Println(vow1("HI There!"))
	fmt.Println(vow1("Why do you ask?"))
	fmt.Println(vow1("Why"))
}