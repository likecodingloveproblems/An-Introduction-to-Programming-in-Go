package main

import (
	"fmt"
)

func main() {
	fib := fibGenerator()
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(5))
}

func fib(i int) int {
	if i <= 1 {
		return i
	}
	return fib(i-2) + fib(i-1)
}

func fibGenerator() func(i int) int {
	var cache = map[int]int{0: 0, 1: 1}
	fib := func(i int) int {
		if _, exist := cache[i]; !exist {
			cache[i] = fib(i-1) + fib(i-2)
		}
		return cache[i]
	}
	return fib
}
