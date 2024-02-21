package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	rs := twoSum(nums, 9)
	fmt.Println(rs)
}

// Solution 1:
//func twoSum(nums []int, target int) []int {
//
//	var rs []int
//
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				rs = append(rs, i, j)
//				return rs
//			}
//		}
//	}
//
//	return rs
//}

// Solution 2
func twoSum(nums []int, target int) []int {

	store := make(map[int]int)

	var rs []int

	for i, v := range nums {
		store[v] = i
	}

	for ix, v := range nums {
		iy, ok := store[target-v]

		if ok && iy != ix {
			rs = []int{ix, iy}
			return rs
		}
	}
	return rs

}
