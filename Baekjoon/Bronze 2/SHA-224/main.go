package main

import (
	"crypto/sha256"
	"fmt"
)
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("%x", sha256.Sum224([]byte(s)))
}