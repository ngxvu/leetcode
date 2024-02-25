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

// Solution 2 - Using hashtable
//func twoSum(nums []int, target int) []int {
//
//	// dùng hashtable để chưa key-value
//	store := make(map[int]int)
//
//	var rs []int
//
//	for i, v := range nums {
//		store[v] = i
//	}
//
//	// lặp qua từng index và value của nums
//	for ix, v := range nums {
//
//		// iy là value của store[target-v]
//		// nếu store[target-v] trùng với key đang chứa trong store
//		// và vì v là value của ix => nên chắc chắn value của ix và iy sẽ bằng target
//		iy, ok := store[target-v]
//
//		if ok && iy != ix {
//			rs = []int{ix, iy}
//			return rs
//		}
//	}
//	return rs
//
//}

// Solution 2 - Using hashtable
func twoSum(nums []int, target int) []int {

	var rs []int

	store := make(map[int]int)

	//[2, 7, 11, 15]
	//store[2] = 0
	//store[7] = 1
	//store[11] = 2
	//store[15] = 3
	//store[9-2] <=> store[7]

	for i, v := range nums {
		store[v] = i
	}

	for i, v := range nums {
		// j là value trong map, store[target-v] sẽ tìm value trong map phù hợp
		j, ok := store[target-v]
		if ok && i != j {
			rs = []int{i, j}
		}
	}

	return rs
}

// Solution 3: Two Pointers
// Two Pointer must sort first
