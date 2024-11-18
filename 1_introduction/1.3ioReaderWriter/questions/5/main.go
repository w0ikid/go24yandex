package main

import (
	"bytes"
	"io"
	"errors"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	if (len(seq) == 0) {
		return false, errors.New("Sequence cannot be empty")
	}


	buffer := make([]byte, 1024)
	remaining := []byte{}

	for {
		// читаем данные из r
		n, err := r.Read(buffer)
		if n > 0 {
			// Конкатенируем остаток и текущий буфер
			data := append(remaining, buffer[:n]...)
			

			if bytes.Contains(data, seq) {
				return true, nil
			}
			
			if len(data) >= len(seq) {
				remaining = data[len(data) - len(seq) + 1:]
			} else {
				remaining = data
			}
		}
		
		if err == io.EOF {
			break
		}

		if err != nil {
			return false, err
		}
	}

	return false, nil
}

func main(){

}