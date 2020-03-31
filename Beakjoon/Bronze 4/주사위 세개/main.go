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
	a=d[0]
	b=d[1]
	c=d[2]
	if a==b && b==c{
		fmt.Println(10000+a*1000)
	} else if a==b||b==c{
		fmt.Println(1000+b*100)
	} else{
		fmt.Println(c*100)
	}
}