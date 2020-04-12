package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)
func main() {
	var s string
	fmt.Scan(&s)
	h := sha1.New()
	io.WriteString(h,s)
	fmt.Printf("%x", h.Sum(nil))
}