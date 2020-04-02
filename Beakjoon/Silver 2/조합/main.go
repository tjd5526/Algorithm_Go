package main

import "fmt"

func main() {
	var n, m float64
	fmt.Scan(&n,&m)
	var f1,f2 float64 = 1,1
	for range make([]int,int(m)){
		f1*=n
		n-=1
		f2*=m
		m-=1
	}
	fmt.Println(f1/f2)
}