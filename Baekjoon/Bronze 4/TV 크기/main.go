package main

import (
	"fmt"
	"math"
)

func main(){
	var a,b,c float64
	fmt.Scanln(&a,&b,&c)
	d := math.Sqrt(math.Pow(a,2)/(math.Pow(b,2)+math.Pow(c,2)))
	fmt.Println(int(b*d), int(c*d))
}