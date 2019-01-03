package main

import (
	"fmt"
)
func pyd1(n int) {
	c := 2 * n - 1
	m := n - 1
	for i := 0; i < n; i ++ {
		str := ""
		for j := 0; j < c; j ++ {
			if j >= m - i && j <= m + i {
				str += "#"
			} else {
				str += " "
			} // ith row: 2*i + 1 "#"
		}
		fmt.Println(str)
	}
}

func pyd2(n, r int, str string) {
	if n == r {
		return
	}
	if len(str) == 2 * n - 1 {
		fmt.Println(str)
		str = ""
		r ++
		pyd2(n, r, str)
	} else if len(str) >= (n-r-1) && len(str) <= (n+r-1) {
		str += "#"
		pyd2(n, r, str)
	} else {
		str += " "
		pyd2(n, r, str)
	}
}

func main() {
	// pyd1(5)
	pyd2(5, 0, "")
}