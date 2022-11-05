package main

import "fmt"

func main() {
	x := 2
	fmt.Println(halve(x))
}

func halve(x int) (half int, isEven bool) {
	half = x / 2
	isEven = x%2 == 0
	return
}
