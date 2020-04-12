package main

import (
	"fmt"
	"strings"
)

func main(){
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	if strings.Repeat(a, len(b))==strings.Repeat(b, len(a)){
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}