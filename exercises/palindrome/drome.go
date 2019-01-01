package main

import "fmt"

func isDrome(str string) bool {
	sls := []rune(str)
	for i := 0; i < len(sls)/2; i++ {
		if sls[i] != sls[len(sls)-1-i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isDrome("abccba"))
	fmt.Println(isDrome("abccbab"))
	fmt.Println(int(005))
}
