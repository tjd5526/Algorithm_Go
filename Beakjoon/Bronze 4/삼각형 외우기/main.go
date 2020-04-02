package main

import "fmt"

func main() {
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	sum := a+b+c
	if a==60 && b==60 && c==60{
		fmt.Println("Equilateral")
	} else if sum==180 && (a==b || b==c || c==a){
		fmt.Println("Isosceles")
	} else if sum==180 && !(a==b || b==c || c==a){
		fmt.Println("Scalene")
	} else{
		fmt.Println("Error")
	}
}