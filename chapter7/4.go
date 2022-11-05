package main

import "fmt"

func main() {
	oddGenerator := makeOddGenerator()
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
}

func makeOddGenerator() func() int {
	num := 1
	generator := func() int {
		temp := num
		num += 2
		return temp
	}
	return generator
}
