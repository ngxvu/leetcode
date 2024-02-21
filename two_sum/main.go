package main

import "fmt"

func main() {
	nums := []int{2, 5, 5, 11}
	rs := twoSum(nums, 10)
	fmt.Println(rs)
}

func twoSum(nums []int, target int) []int {

	var rs []int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				rs = append(rs, i, j)
				return rs
			}
		}
	}

	return rs
}
