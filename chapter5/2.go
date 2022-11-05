package main

import "fmt"

func main() {
	maxNum := 100
	dividableNum := 3
	for num := 1; num < maxNum; num++ {
		if num%dividableNum == 0 {
			fmt.Println(num)
		}
	}
}
