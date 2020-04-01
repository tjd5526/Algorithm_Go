package main

import (
	"fmt"
	"math"
)

func main(){
	var A,B,C,D,E float64
	fmt.Scan(&A,&B,&C,&D,&E)
	fmt.Println(math.Min(math.Min(A,B),C)+math.Min(D,E)-50)
}