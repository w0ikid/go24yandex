package main

import (
	"fmt"
	// "strings"
)

// func Permutations(input string) []string {
// 	length := len([]rune(input))
// 	n := factorial(length)
// 	/*
// 	123 -> 132 231
// 	*/
// }

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main(){
	fmt.Println("somethign")
}