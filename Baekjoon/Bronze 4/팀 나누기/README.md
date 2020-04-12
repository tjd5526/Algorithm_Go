[팀 나누기](https://www.acmicpc.net/problem/13866)
```go
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
```