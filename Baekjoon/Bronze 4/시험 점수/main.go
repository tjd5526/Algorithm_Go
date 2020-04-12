package main

import (
	"fmt"
	"math"
)

func main(){
	var a,b,c,d,e,f,g,h float64
	fmt.Scan(&a,&b,&c,&d,&e,&f,&g,&h)
	fmt.Println(math.Max(a+b+c+d,e+f+g+h))
}