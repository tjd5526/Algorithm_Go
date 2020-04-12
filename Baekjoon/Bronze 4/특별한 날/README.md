[특별한 날](https://www.acmicpc.net/problem/10768)
```go
package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a,&b)
	b=a*30+b
	c:=78
	if b<c{
		fmt.Println("Before")
	} else if b>c{
		fmt.Println("After")
	} else if b==c{
		fmt.Println("Special")
	}
}
```