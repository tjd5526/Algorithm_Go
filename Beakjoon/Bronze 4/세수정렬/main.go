package main

import (
	"fmt"
	"sort"
)
func main(){
	var a,b,c int
	fmt.Scanln(&a,&b,&c)
	d:=[]int{a,b,c}
	sort.Ints(d)
	fmt.Println(d[0], d[1], d[2])
}