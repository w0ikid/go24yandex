package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	switch {
	case n > 0:
		fmt.Println("Число положительное")
	case n < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Введен ноль")
	}
}