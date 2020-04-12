[SHA-384](boj.kr/10931)
```go
package main

import (
	"crypto/sha512"
	"fmt"
)
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("%x", sha512.Sum384([]byte(s)))
}
```