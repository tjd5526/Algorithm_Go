[사분면 고르기](https://www.acmicpc.net/problem/14681)
```go
package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a,&b)
	if a>0{
		if b>0{
			fmt.Print(1)
		} else{
			fmt.Print(4)
		}
	} else{
		if b>0{
			fmt.Print(2)
		} else{
			fmt.Print(3)
		}
	}
}
```