[과자](https://www.acmicpc.net/problem/10156)
```go
package main

import "fmt"

func main() {
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	if a*b<c{
		fmt.Println(0)
	} else{
		fmt.Println(a*b-c)
	}
}
```