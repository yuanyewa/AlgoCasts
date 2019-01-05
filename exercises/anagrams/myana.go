package main

import (
	"fmt"
	"time"
	"unicode"
)

func ana1(str1, str2 string) bool {
	r1 := []rune(str1)
	r2 := []rune(str2)
	m := make(map[rune]int)
	for _, r := range r1 {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			m[unicode.ToLower(r)]++
		}
	}
	for _, r := range r2 {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			if m[unicode.ToLower(r)] == 0 {
				return false
			}
			m[unicode.ToLower(r)]--
			if m[unicode.ToLower(r)] == 0 {
				delete(m, unicode.ToLower(r))
			}
		}
	}
	return len(m) == 0
	// for _, v := range m {
	// 	if v > 0 {
	// 		return false
	// 	}
	// }
	// return true
}

func main() {
	fmt.Println(ana1("abc dd d", "d dDB ca"))
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			ana1("abc dd d abcde", "d dDB caabcde")
		}
		fmt.Println(time.Since(start))
	}
}
