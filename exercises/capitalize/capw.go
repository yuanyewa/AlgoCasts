package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func cap1(str string) string {
	sls := strings.Fields(str)
	for i := range sls {
		if unicode.IsLower(rune(sls[i][0])) {
			sls[i] = string(unicode.ToUpper(rune(sls[i][0]))) + sls[i][1:]
		}
	}
	return strings.Join(sls, " ")
}
func cap2(str string) string {
	sls := strings.Fields(str)
	for i := range sls {
		sls[i] = string(unicode.ToUpper(rune(sls[i][0]))) + sls[i][1:]
	}
	return strings.Join(sls, " ")
}

func cap3(str string) string {
	out := string(unicode.ToUpper(rune(str[0])))
	for i := 1; i < len(str); i++ {
		if str[i-1] == byte(' ') {
			out += string(unicode.ToUpper(rune(str[i])))
		} else {
			out += string(str[i])
		}
	}
	return out
}
func main() {
	str := "this is me with good"
	fmt.Println(cap1(str))
	fmt.Println(strings.Title(str))
	fmt.Println(cap3(str))
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			cap1(str)
		}
		fmt.Println(time.Since(start))
	}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			strings.Title(str)
		}
		fmt.Println(time.Since(start))
	}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			cap3(str)
		}
		fmt.Println(time.Since(start))
	}
}
