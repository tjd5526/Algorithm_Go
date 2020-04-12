[곱셈](https://www.acmicpc.net/problem/2588)
```go
package main

import (
	"fmt"
)

func main(){
	var a,b int
	fmt.Scan(&a,&b)
	fmt.Println(a*(b%10))
	fmt.Println(a*(b%100/10))
	fmt.Println(a*(b/100))
	fmt.Println(a*b)
}
```