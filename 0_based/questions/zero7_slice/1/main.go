package main

import (
	"fmt"
	"errors"
)

func UnderLimit1(nums []int, limit int, n int) ([]int, error){
	a := make([]int, 0, n)
	for i := 0; i < len(nums); i++ {
		if nums[i] < limit && n != 0{
			a = append(a, nums[i])
			n--
		}
	}

	return a, nil
}

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("n должно быть больше 0")
	}

	result := make([]int, 0, n)
	
	for _, num := range nums {
		if num < limit {
			result = append(result, num)
			n--
		}
		if n == 0 {
			break
		}
	}

	return result, nil
}

func main(){
	a := []int{1,2,3,4,5}
	fmt.Println(UnderLimit(a, 4, 2))
}