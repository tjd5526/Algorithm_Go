package main

import (
	"fmt"
	"sort"
)

func main() {
	var a,b,c,d,e,f int
	ans := 0
	fmt.Scan(&a,&b,&c,&d,&e,&f)
	g:=[]int{a,b,c,d}
	s:=[]int{e,f}
	sort.Ints(g)
	sort.Ints(s)
	g= g[1:]
	for _, j := range g{
		ans+=j
	}
	fmt.Print(ans+s[1])
}