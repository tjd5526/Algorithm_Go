package main

import "fmt"

func main() {
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)

	if b+c<a+d{
		fmt.Print(b+c)
	} else{
		fmt.Print(a+d)
	}
}