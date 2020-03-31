[사파리월드](https://www.acmicpc.net/problem/2420)
```go
package main

import (
	"fmt"
	"math"
)

func main(){
	var a, b float64
	fmt.Scanln(&a,&b)
	fmt.Println(int(math.Abs(a-b)))
}
```