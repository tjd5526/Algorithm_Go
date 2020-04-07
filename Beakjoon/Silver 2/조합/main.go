package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n, m float64
	fmt.Scan(&n,&m)
	var fact = new(big.Int)
	fact.MulRange(1,100)
	str := fact.String()

}