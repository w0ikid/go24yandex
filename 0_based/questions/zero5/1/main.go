package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	// Вычисление дискриминанта
	Discriminant := b*b - 4*a*c

	if Discriminant < 0 {
		// Если дискриминант меньше нуля, корней нет
		fmt.Println("0 0")
	} else if Discriminant == 0 {
		// Если дискриминант равен нулю, один корень
		x := -b / (2 * a)
		fmt.Printf("%.4f\n", x)
	} else {
		// Если дискриминант больше нуля, два корня
		x1 := (-b - math.Sqrt(Discriminant)) / (2 * a)
		x2 := (-b + math.Sqrt(Discriminant)) / (2 * a)

		// Выводим корни в порядке возрастания
		if x1 < x2 {
			fmt.Printf("%.4f %.4f\n", x1, x2)
		} else {
			fmt.Printf("%.4f %.4f\n", x2, x1)
		}
	}
}


func main(){
	SqRoots()
}