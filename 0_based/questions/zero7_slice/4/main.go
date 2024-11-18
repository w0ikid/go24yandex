package main

import "fmt"

func Join(nums1, nums2 []int) []int {
	newNums := make([]int, 0, len(nums1) + len(nums2))
	newNums = append(newNums, nums1...)
	newNums = append(newNums, nums2...)
	return newNums
}

func main(){
	newTest1 := []int{1, 2, 3}
	newTest2 := []int{4, 5, 6, 7}
	
	fmt.Println(Join(newTest1, newTest2))
	fmt.Println(cap(Join(newTest1, newTest2)))
	
	fmt.Println(cap(newTest1))
	fmt.Println(cap(newTest2))

}