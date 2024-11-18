package main

import (
	"fmt"
	"strings"
)

type UpperWriter struct {
	UpperString string
}

func (u *UpperWriter) Write(p []byte) (n int, err error){
	input := string(p)

	upper := strings.ToUpper(input)

	u.UpperString = upper

	return len(p), nil
}

func main(){
	writer := &UpperWriter{}

	// input value

	input := []byte("some text")

	_, err := writer.Write(input)

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Result:", writer.UpperString)
	}
}