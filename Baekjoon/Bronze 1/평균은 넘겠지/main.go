package main

import "fmt"

func main(){
	var c,d int
	fmt.Scan(&c)
	d=c
	for i:=0;i<d;i++{
		fmt.Scan(&c)
		var total float64 = 0
		count := 0
		d := make([]float64, c)
		for i, _ := range d{
			fmt.Scan(&d[i])
			total+=d[i]
		}
		total/=float64(c)
		for _, j := range d{
			if j>total{
				count++
			}
		}
		fmt.Printf("%.3f%%\n", float64(count)/float64(c)*100)
	}
}
