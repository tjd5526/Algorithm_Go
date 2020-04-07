package main

import "fmt"

func main() {
	var ans, cot int = 0,0
	a := make([]int, 5)
	fmt.Scan(&ans)
	for i, _ := range a{
		fmt.Scan(&a[i])
	}
	for _, j := range a{
		if ans==j{
			cot+=1
		}
	}
	fmt.Println(cot)
}