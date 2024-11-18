package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	count := 0
	for _, char := range input {
		if string(char) == "🔑" {
			count++
		}
	}

	fmt.Printf("%d🔑\n", count)
}
