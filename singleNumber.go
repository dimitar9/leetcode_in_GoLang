package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 4, 5, 6, 4, 5}
	fmt.Println((singleNumber(nums, 8)))
}
func singleNumber(nums []int, n int) (ret int) {

	for i := 1; i < n; i++ {
		nums[0] = nums[0] ^ nums[i]
	}
	return nums[0]
}
