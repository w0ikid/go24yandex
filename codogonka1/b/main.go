package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inventory := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		key := strings.TrimSpace(scanner.Text())
		if _, exists := inventory[key]; exists {
			fmt.Printf("ключ %s улучшен\n", key)
		} else {
			inventory[key] = true
			fmt.Printf("ключ %s добавлен в инвентарь\n", key)
		}
	}
}
