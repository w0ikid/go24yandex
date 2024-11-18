package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var coins []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if value, err := strconv.Atoi(line); err == nil {
			coins = append(coins, value)
		} else if line == "сумма" {
			totalCoins := len(coins)
			totalValue := 0
			for _, coin := range coins {
				totalValue += coin
			}
			fmt.Printf("Всего монет: %d. Общий номинал: %d.\n", totalCoins, totalValue)
			break
		}
	}
}
