package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalCost := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "поступило" {
			break
		}
		if cost, err := strconv.Atoi(line); err == nil {
			totalCost += cost
		}
	}

	if scanner.Scan() {
		payment, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
			return
		}

		if payment > totalCost {
			fmt.Println(payment - totalCost)
		} else if payment < totalCost {
			fmt.Printf("доплатите %d\n", totalCost-payment)
		} else {
			fmt.Println("спасибо, что без сдачи")
		}
	}
}
