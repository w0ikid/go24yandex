package main

import (
	"fmt"
)

func Add(a, b float64) float64 {
	return a + b
}

func Multiply(a, b float64) float64{
	return a * b
}

func PrintNumbersAscending(n int){
	for i := 1; i <= n; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
}

func main(){

}