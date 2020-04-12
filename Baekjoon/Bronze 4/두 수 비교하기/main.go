package main

import "fmt"

func main(){
	var a, b int
	fmt.Scanln(&a, &b)
	if a<b{
		fmt.Println("<")
	}
	if a==b{
		fmt.Println("==")
	}
	if a>b{
		fmt.Println(">")
	}
}