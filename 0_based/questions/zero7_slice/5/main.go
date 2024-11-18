package main

import "fmt"

// 

func Mix(nums []int) []int {
	n := len(nums) / 2
	mixedSlice := make([]int, len(nums))

	for i := 0; i < n; i++ {
		mixedSlice[2*i] = nums[i]
		mixedSlice[2*i+1] = nums[n+i]
	}

	return mixedSlice
}

func main(){
	a := []int{1,2,3,4,5,6}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 6

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a[len(a)/2])

	fmt.Println(Mix(a))

}