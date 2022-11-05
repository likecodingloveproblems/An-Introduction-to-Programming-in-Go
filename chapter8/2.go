package main

import "fmt"

func main() {
	num := new(int)
	*num = 4
	square(num)
	fmt.Println(*num)
}

func square(num *int) {
	*num = *num * *num
}
