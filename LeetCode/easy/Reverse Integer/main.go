package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var ans int
	for x != 0 {
		ans = ans*10 + x%10
		if ans < math.MinInt32 || ans > math.MaxInt32 {
			return 0
		}
		x /= 10
	}
	return ans
}

func main() {
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}
