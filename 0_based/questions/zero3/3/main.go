package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Число равно нулю")
		return
	}

	zero := "отрицательное"
	if n > 0 {
		zero = "положительное"
	}

	even := "нечетное"
	if n%2 == 0 {
		even = "четное"
	}

	fmt.Printf("Число %s и %s\n", zero, even)
}
