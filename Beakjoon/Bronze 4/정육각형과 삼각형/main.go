package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Scan(&a)
	b := math.Pow(a,2)
	fmt.Println(b/4*math.Sqrt(3))
}