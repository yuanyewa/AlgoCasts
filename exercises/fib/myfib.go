package main

import "fmt"

func fib1(n int) int {
	if n == 0 {
		return 0
	}
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}

var cached = map[int]int{}

func fib3(n int) int {
	if n < 2 {
		return n
	}
	if v, ok := cached[n]; ok {
		return v
	}
	v := fib2(n-1) + fib2(n-2)
	cached[n] = v
	return v
}

func main() {
	for n := 1; n < 11; n++ {
		fmt.Println(fib3(n))
	}
	fmt.Println(fib1(10))

}
