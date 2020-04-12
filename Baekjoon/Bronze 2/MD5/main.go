package main

import (
	"crypto/md5"
	"fmt"
	"io"
)
func main() {
	var s string
	h := md5.New()
	fmt.Scan(&s)

	io.WriteString(h, s)
	fmt.Printf("%x", h.Sum(nil))
}