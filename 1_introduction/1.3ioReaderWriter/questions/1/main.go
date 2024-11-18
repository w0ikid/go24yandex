package main

import (
	"fmt"
	"io"
)

func WriteString(s string, w io.Writer) error {
    _, err := w.Write([]byte(s))
    if err != nil {
        return fmt.Errorf("please try write something else")
    }
    return nil
}

func main(){
    
}