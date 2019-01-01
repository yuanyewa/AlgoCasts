package main

import (
	"fmt"
	"time"
)

func rev1(str string) string {
	// r := strings.Split(str, "")
	r := []rune(str)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	// return strings.Join(r, "")
	return string(r)
}
func rev2(str string) string {
	c := []rune(str)
	r := make([]rune, 0, len(c))
	for _, t := range c {
		r = append([]rune{t}, r...)
	}
	return string(r)
}
func main() {
	{
		fmt.Println(rev1("hello世界"))
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			rev1("hello世界")
		}
		fmt.Println(time.Since(start))
	}
	{
		fmt.Println(rev2("hello世界"))
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			rev2("hello世界")
		}
		fmt.Println(time.Since(start))
	}
}
