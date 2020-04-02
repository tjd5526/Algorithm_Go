[전자레인지](https://www.acmicpc.net/problem/10162)
```go
package main

import "fmt"

func main() {
	var a, re1,re2,re3 int
	A,B,C := 300,60,10
	fmt.Scan(&a)
	if a%C!=0{
		fmt.Println(-1)
	} else{
		re1, a = a/A, a%A
		re2, a = a/B, a%B
		re3, a = a/C, a%C
		fmt.Println(re1,re2,re3)
	}
}
```