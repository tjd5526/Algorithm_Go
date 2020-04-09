package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	a := strconv.Itoa(x)
	b := ""
	for _, j := range a {
		b = string(j) + b
	}
	return a == b
}

func main() {
	fmt.Println(isPalindrome(121))
}
