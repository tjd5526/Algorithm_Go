[윤년](https://www.acmicpc.net/problem/2753)
```go
package main

import "fmt"

func main(){
	var a int
	fmt.Scan(&a)
	if a%4==0 && (a%100!=0 || a%400==0){
		fmt.Println(1)
	} else{
		fmt.Println(0)
	}
}
```