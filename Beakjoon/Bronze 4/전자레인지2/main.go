package main

import "fmt"

func main() {
	var a,b,c,d,e int
	fmt.Scan(&a,&b,&c,&d,&e)
	if a<0{
		fmt.Print(-a*c+d+b*e)
	} else{
		fmt.Print((b-a)*e)
	}
}