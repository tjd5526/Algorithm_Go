[두 수 비교하기](https://www.acmicpc.net/problem/1330)
```go
package main

import "fmt"

func main(){
	var a, b int
	fmt.Scanln(&a, &b)
	if a<b{
		fmt.Println("<")
	}
	if a==b{
		fmt.Println("==")
	}
	if a>b{
		fmt.Println(">")
	}
}
```