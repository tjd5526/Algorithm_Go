[무한문자열](https://www.acmicpc.net/problem/12871)
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	if strings.Repeat(a, len(b))==strings.Repeat(b, len(a)){
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
```
strings.Compare로 비교하는 방법도 있지만 godoc에선 이것보단 ==이나 >,< 같은 연산자로 비교하는걸 추천하고 있다.