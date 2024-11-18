package main

import (
	"fmt"
	"io"
)

type MyStruct struct {}

func (m *MyStruct) Read(p []byte) (n int, err error)

type SomeNewStruct struct{
	Name string
	id int
}

func (s *SomeNewStruct) Error() string {
	return fmt.Sprintf("Error %s, %d", s.Name, s.id)
}

var _ io.Reader = (*MyStruct)(nil)

var _ error = (*SomeNewStruct)(nil)

func main(){
	
}


