package main

import (
	"fmt"
	"strings"
)

func removeSpace(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

func IsPalindrome(input string) bool {
	newString := strings.ToLower(removeSpace(input))
	
	n := len([]rune(newString)) // 5
	for i := 0; i < n/2; i++ {
		if []rune(newString)[i] != []rune(newString)[n-1-i]{
			return false
		}
	}
	return true
}

func main(){
	sometext := " a b c       cba  "
	n := len([]rune(sometext))
	fmt.Println(n)
	fmt.Println(IsPalindrome(sometext))
}