package main

import (
	"fmt"
	"time"
)

func step1(n int) {
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			if j <= i {
				str += "#"
			} else {
				str += " "
			}
		}
		// fmt.Printf("%q\n", str)
	}
}

func step2(n, r int, str string) {
	if r == n {
		return
	}
	if len(str) >= n {
		// fmt.Println(str)
		step2(n, r+1, "")
	} else if len(str) <= r {
		str += "#"
		step2(n, r, str)
	} else {
		str += " "
		step2(n, r, str)
	}
}
func main() {

	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			step1(10)
		}
		fmt.Println(time.Since(start))
	}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			step2(10, 0, "")
		}
		fmt.Println(time.Since(start))
	}
}
