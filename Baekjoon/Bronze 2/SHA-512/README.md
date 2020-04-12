[SHA-512](boj.kr/10932)
```go
package main

import (
	"crypto/sha512"
	"fmt"
	"io"
)
func main() {
	var s string
	h := sha512.New()
	fmt.Scan(&s)

	io.WriteString(h, s)
	fmt.Printf("%x", h.Sum(nil))
}
```