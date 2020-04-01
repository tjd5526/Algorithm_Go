[체스판 조각](https://www.acmicpc.net/problem/3004)
```go
package main

import "fmt"

func main(){
	var a int
	fmt.Scan(&a)
	fmt.Println((a/2 + 1)*(a - a/2 + 1))
}
```