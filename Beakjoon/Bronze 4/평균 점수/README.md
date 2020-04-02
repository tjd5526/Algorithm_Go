[평균 점수](https://www.acmicpc.net/problem/10039)
```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	sum := 0
	fmt.Scan(&a[0], &a[1], &a[2], &a[3], &a[4])

	for _, i := range a {
		if i < 40 {
			i = 40
		}
		sum += i
	}
	fmt.Println(sum/5)
}
```