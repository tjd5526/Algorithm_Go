[시험 성적](https://www.acmicpc.net/problem/9498)
```go
package main

import "fmt"

func main(){
	var a int
	fmt.Scan(&a)
	if a>=90{
		fmt.Println("A")
	} else if a>=80{
		fmt.Println("B")
	} else if a>=70{
		fmt.Println("C")
	} else if a>=60{
		fmt.Println("D")
	} else{
		fmt.Println("F")
	}
}
```