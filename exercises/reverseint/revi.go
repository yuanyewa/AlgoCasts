package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func rev1(p int, n int) (int, int) {
	if n < 10 && n > -10 {
		return 0, p*10 + n
	}
	return rev1(p*10+n%10, n/10)
}

type mystring string

func (s mystring) reverse() mystring {
	sls := []rune(s)
	for i := 0; i < len(sls)/2; i++ {
		sls[i], sls[len(sls)-1-i] = sls[len(sls)-1-i], sls[i]
	}
	return mystring(sls)
}
func rev2(n int) int {
	m, _ := strconv.Atoi(string(mystring(strconv.Itoa(int(math.Abs(float64(n))))).reverse()))
	if n < 0 {
		return -m
	}
	return m
}
func main() {
	fmt.Println(rev1(0, 123))
	fmt.Println(rev1(0, -98765432))
	fmt.Println(-24 % 10)
	fmt.Println(rev2(-987654321))
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			rev1(0, -987654321)
		}
		fmt.Println(time.Since(start))
	}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			rev2(-987654321)
		}
		fmt.Println(time.Since(start))
	}
}
