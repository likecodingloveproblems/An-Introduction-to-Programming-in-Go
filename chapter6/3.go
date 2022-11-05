package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(Min(x))
}

func Min(nums []int) (min int) {
	min = nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return
}
