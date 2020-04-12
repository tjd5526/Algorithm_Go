[수도요금](https://www.acmicpc.net/problem/10707)
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var X, Y, Yl, Yo, U int
	fmt.Scan(&X, &Y, &Yl, &Yo, &U)
	X = X*U
	if Yl<U{
		Y+=(U-Yl)*Yo
	}
	fmt.Println(math.Min(float64(X),float64(Y)))
}
```