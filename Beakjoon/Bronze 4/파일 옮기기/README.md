[파일 옮기기](https://www.acmicpc.net/problem/11943)
```go
package main

import "fmt"

func main() {
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)
	if b+c<a+d{
		fmt.Print(b+c)
	} else{
		fmt.Print(a+d)
	}
}
```