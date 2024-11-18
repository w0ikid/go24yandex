package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nfact := 1
	for i := 1; i <= n; i++ {
		nfact = nfact * i
	}
	fmt.Println(nfact)
}
