[오븐시계](https://www.acmicpc.net/problem/2525)
```go
package main

import "fmt"

func main(){
	var a, b, c int
	fmt.Scanln(&a,&b)
	fmt.Scanln(&c)
	b+=c
	if b>59{
		a+=b/60
		b%=60
		a%=24
		fmt.Println(a,b)
	} else{
		fmt.Println(a,b)
	}
}
```