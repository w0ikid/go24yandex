package main

import (
	"fmt"
)

func IsPowerOfTwoRecursive(N int) {
	if N == 0 {
		fmt.Println("NO")
		return
	}
	if N == 1 {
		fmt.Println("YES")
		return
	}
	if N % 2 != 0 {
		fmt.Println("NO")
		return
	}
	IsPowerOfTwoRecursive(N / 2)
}


func main(){
	var a int
	var b int
	a = 5
	b = 3
	fmt.Println(a/b)
}
