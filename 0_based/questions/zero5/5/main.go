package main

import(
	"fmt"
)

func CalculateSeriesSum(n int) float64 {
	if n == 0 {
		return 0
	}
	return 1/float64(n) + CalculateSeriesSum(n-1)
}

func main(){
	var a int
	fmt.Scan(&a)
	fmt.Println(CalculateSeriesSum(a))
}