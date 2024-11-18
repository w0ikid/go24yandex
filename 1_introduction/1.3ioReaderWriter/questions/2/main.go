package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadString(r io.Reader) (string, error) {
	var builder strings.Builder
	buffer := make([]byte, 1024)
	for {
		n, err := r.Read(buffer)
		if n > 0 {
			builder.Write(buffer[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("read error: %w", err)
		}
	}
	return builder.String(), nil
}