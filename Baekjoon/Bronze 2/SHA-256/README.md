[SHA-256](boj.kr/10930)
```go
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
)
func main() {
	var s string
	fmt.Scan(&s)
	h := sha256.New()
	io.WriteString(h,s)
	fmt.Printf("%x", h.Sum(nil))
}
```