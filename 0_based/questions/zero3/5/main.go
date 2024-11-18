package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		fmt.Println("Все числа равны")
	} else if a == b || a == c || b == c {
		fmt.Println("Два числа равны")
	} else {
		fmt.Println("Все числа разные")
	}
}
