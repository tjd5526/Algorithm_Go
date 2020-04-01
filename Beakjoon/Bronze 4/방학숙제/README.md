[방학숙제](https://www.acmicpc.net/problem/5532)
```go
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
```