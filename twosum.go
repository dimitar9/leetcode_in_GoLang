package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15, 18, 22}
	fmt.Println((twoSum(nums, 37)))
}
func twoSum(nums []int, target int) (index1, index2 int) {

	hm := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hm[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := hm[target-nums[i]]; ok {
			index1 = i + 1
			index2 = hm[target-nums[i]] + 1
			return index1, index2
		}
	}
	return
}
