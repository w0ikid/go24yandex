package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)

	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}

	if a == b && b == c {
		fmt.Println("Все числа равны")
	} else if a == b || a == c || b == c {
		fmt.Println("Два числа равны")
	} else {
		fmt.Println("Все числа разные")
	}
}
