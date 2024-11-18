package main

import "fmt"

func SliceCopy(nums []int) []int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	return numsCopy
}