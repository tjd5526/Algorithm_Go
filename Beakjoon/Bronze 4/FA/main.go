package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Scan(&a)
	var ans float64 = 0
	for {
		if ans==a{
			fmt.Println("FA")
		}else{
			ans = len(string(a))
			ans = ans*a/math.Pow(10, ans)
		}
	}

}