package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if b < c {
		b, c = c, b
	}

	minCoins := -1

	for i := 0; i*b <= a; i++ {
		remaining := a - i*b
		if remaining%c == 0 {
			totalCoins := i + remaining/c
			if minCoins == -1 || totalCoins < minCoins {
				minCoins = totalCoins
			}
		}
	}

	if minCoins == -1 {
		fmt.Println("переводом, пожалуйста")
	} else {
		fmt.Println(minCoins)
	}
}
