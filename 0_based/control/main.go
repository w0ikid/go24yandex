package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	var nums []float64
	var ops []rune

	applyOperation := func() error {
		if len(nums) < 2 {
			return errors.New("недостаточно операндов")
		}
		if len(ops) == 0 {
			return errors.New("недостаточно операторов")
		}

		b := nums[len(nums)-1]
		a := nums[len(nums)-2]
		op := ops[len(ops)-1]

		nums = nums[:len(nums)-2]
		ops = ops[:len(ops)-1]

		var res float64
		switch op {
		case '+':
			res = a + b
		case '-':
			res = a - b
		case '*':
			res = a * b
		case '/':
			if b == 0 {
				return errors.New("деление на ноль")
			}
			res = a / b
		default:
			return errors.New("неизвестная операция")
		}

		nums = append(nums, res)
		return nil
	}

	precedence := func(op rune) int {
		switch op {
		case '+', '-':
			return 1
		case '*', '/':
			return 2
		default:
			return 0
		}
	}

	parseNumber := func(i *int) (float64, error) {
		start := *i
		for *i < len(expression) && (unicode.IsDigit(rune(expression[*i])) || expression[*i] == '.') {
			*i++
		}
		numStr := expression[start:*i]
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, errors.New("некорректное число: " + numStr)
		}
		return num, nil
	}

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		if unicode.IsDigit(ch) || ch == '.' {
			num, err := parseNumber(&i)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
			i--
		} else if ch == '(' {
			ops = append(ops, ch)
		} else if ch == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if err := applyOperation(); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 || ops[len(ops)-1] != '(' {
				return 0, errors.New("несоответствие скобок")
			}
			ops = ops[:len(ops)-1]
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(ch) {
				if err := applyOperation(); err != nil {
					return 0, err
				}
			}
			ops = append(ops, ch)
		} else {
			return 0, errors.New("некорректный символ: " + string(ch))
		}
	}

	for len(ops) > 0 {
		if err := applyOperation(); err != nil {
			return 0, err
		}
	}

	if len(nums) != 1 {
		return 0, errors.New("некорректное выражение")
	}

	return nums[0], nil
}

func main() {
	expr := "3 + (2 * 2) - 1 / 2"
	result, err := Calc(expr)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
