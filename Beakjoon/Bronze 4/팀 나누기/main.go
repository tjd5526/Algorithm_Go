package main

import (
	"fmt"
	"math"
)

func main() {
	var a,b,c,d float64
	fmt.Scan(&a,&b,&c,&d)
	fmt.Print(math.Abs(a+d-b-c))
}