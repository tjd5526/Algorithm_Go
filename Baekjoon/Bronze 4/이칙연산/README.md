[이칙연산](https://www.acmicpc.net/problem/15726)
```go
package main

import "fmt"

func main() {
	var a,b,c float64
	fmt.Scan(&a,&b,&c)
	d:=a*b/c
	e:=a/b*c
	if d>e{
		fmt.Println(int(d))
	} else{
		fmt.Println(int(e))
	}
}
```