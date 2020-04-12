package main

import (
	"fmt"
	"math"
)

func main(){
	var L,A,B,C,D float64
	fmt.Scan(&L,&A,&B,&C,&D)
	fmt.Println(L-math.Max(math.Ceil(A/C), math.Ceil(B/D)))
}