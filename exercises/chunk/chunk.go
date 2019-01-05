package main

import "fmt"

func chunk(sls []interface{}, size int) [][]interface{} {
	lens := len(sls)
	newLens := (lens-1)/size + 1
	dsls := make([][]interface{}, 0, newLens)
	for i := 0; i < newLens; i++ {
		start := i * size
		end := (i + 1) * size
		if end > lens {
			end = lens
		}
		dsls = append(dsls, sls[start:end])
	}
	return dsls
}
func main() {
	fmt.Println(chunk([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2))
	fmt.Println(chunk([]interface{}{1, 2, "xxxx", "a", []int{1, 2, 3}, 6, 7, 8, 9}, 10))
}
