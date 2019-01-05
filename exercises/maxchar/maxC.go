package main

import "fmt"

func maxc(str string) string {
	m := make(map[rune]int)
	var mc rune
	var mi int
	for _, r := range []rune(str) {
		m[r]++
		if m[r] > mi {
			mi = m[r]
			mc = r
		}
	}
	return string(mc)
}

func main() {
	fmt.Println(maxc("abbbbcccd"))
	x := make([][]int, 0, 3)
	x = append(x, []int{1, 2, 3})
	x = append(x, []int{1, 2, 3})
	fmt.Println(x)
}
