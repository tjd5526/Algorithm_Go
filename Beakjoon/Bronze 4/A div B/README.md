```go
package main

import "fmt"

func main(){
	var a, b float64
	fmt.Scanln(&a, &b)
	fmt.Println(a/b)
}
```
첫번째 코드다.  
828KB 4ms 98B
```go
package main

import "fmt"

func main(){
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(a/b)
}
```
두번째 코드다.  
Scanln보다 메모리가 더 나갔다.
836KB 4ms 96B