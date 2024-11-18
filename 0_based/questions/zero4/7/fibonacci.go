package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 || n == 3 {
		return 1
	}
	sum := fibonacci(n-1) + fibonacci(n-2)
	return sum
}

func main() {
	var n int
	// var nearFib int
	fmt.Scan(&n)
	k := 1
	for true {
		if fibonacci(k) >= n {
			// nearFib = fibonacci(i)
			break
		}
		k++
	}
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(k))
		k++
	}
}
