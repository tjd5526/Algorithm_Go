package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	var f1 float64 = 1
	for range make([]int,int(n)){
		f1*=n
		n-=1
	}
	fmt.Println(f1)
}