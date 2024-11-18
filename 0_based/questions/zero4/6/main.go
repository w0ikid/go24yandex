package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	// if n < 0 {
	// 	fmt.Println("Некорректный ввод")
	// 	return
	// }
	sum := 0
	for i := 1; i <= n; i++ {
		if i % 3 != 0 && i % 5 != 0{
			sum = sum + i
		}
	}
	fmt.Println(sum)
}
