package main

import "fmt"

func main(){
	var a int
	fmt.Scan(&a)
	fmt.Println((a/2 + 1)*(a - a/2 + 1))
}