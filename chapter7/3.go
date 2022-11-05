package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, -5}
	fmt.Println(max(x...))
}

func max(nums ...int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}
