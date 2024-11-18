package main

import "fmt"

func isLatin(input string) bool {
	for _, letter := range input {
		if !((letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122)) {
			return false
		}
	}
	return true
}

func main(){
	textEng := "aasdfasdf" // A 65 Z 90 a 97 z 122
	textRu := "фвыафыафывңіа"
	firstletterEng := []rune(textEng)[0]
	firstletterRu := []rune(textRu)[0]
	
	fmt.Println(firstletterEng)
	fmt.Println(firstletterRu)

	fmt.Println(isLatin(textEng))
	fmt.Println(isLatin(textRu))
}