package main

import "fmt"

func main(){
	weather := "rain"

	switch weather {
	case "rain":
		fmt.Println(";<")
	case "sunny":
		fmt.Println(":>")
	default:
		fmt.Println("Weather hz")
	}

}