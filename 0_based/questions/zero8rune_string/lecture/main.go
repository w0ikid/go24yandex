package main

import (
	"fmt"
)

// PolynomialRollingHash вычисляет хэш строки s по полиномиальному методу
func PolynomialRollingHash(s string, p, m int) int {
	hash := 0
	pPow := 1

	for i := 0; i < len(s); i++ {
		hash = (hash + int(s[i])*pPow) % m
		pPow = (pPow * p) % m
	}

	return hash
}

func main() {
	s1 := "123123123123341312341234231"
	s2 := "123123123112123123123123123" // "soldat"
	p := 128       // Выбранное значение для p
	m := 1_000_000_7 // Выбранное значение для m (обычно большое простое число)

	hashValue1 := PolynomialRollingHash(s1, p, m)
	hashValue2 := PolynomialRollingHash(s2, p, m)
	fmt.Printf("Hash of '%s' is %d\n", s1, hashValue1)
	fmt.Printf("Hash of '%s' is %d\n", s2, hashValue2)
}
