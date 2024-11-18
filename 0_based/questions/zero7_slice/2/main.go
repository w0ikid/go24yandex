package main

import "fmt"

func Clean(nums []int, x int) ([]int){
	for i := 0; i < len(nums); {
		if nums[i] == x {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}	
		i++
	}
	return nums
}

func main(){
	nums := []int{1,2,3,4,5}
	x := 6

	newnums := Clean(nums, x)
	fmt.Println(newnums)
}
// 
// 1 2 2 3 2 4 2
// 1 2 3 2 4 2
// 1 3 2 4 2
// 1 3 4 2
