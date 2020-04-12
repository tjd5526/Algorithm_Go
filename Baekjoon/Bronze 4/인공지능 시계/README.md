[인공지능시계](https://www.acmicpc.net/problem/2530)
```go
package main

import "fmt"

func main(){
	var a, b, c, d int
	fmt.Scanln(&a,&b,&c)
	fmt.Scanln(&d)
	c+=d
	b+=c/60
	a+=b/60
	fmt.Println(a%24,b%60,c%60)
}
```