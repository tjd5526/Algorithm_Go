[주사위 세개](https://www.acmicpc.net/problem/2480)
```go
package main

import (
	"fmt"
	"sort"
)

func main(){
	var a,b,c int
	fmt.Scanln(&a,&b,&c)
	d:=[]int{a,b,c}
	sort.Ints(d)
	a=d[0]
	b=d[1]
	c=d[2]
	if a==b && b==c{
		fmt.Println(10000+a*1000)
	} else if a==b||b==c{
		fmt.Println(1000+b*100)
	} else{
		fmt.Println(c*100)
	}
}
```
정렬을 해주면 만약에 같은 숫자가 있을 경우 무조건 b에서 만나기 때문에 이렇게 짰다.