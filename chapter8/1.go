package main

import "fmt"

func main() {
	x := 5
	incr(&x)
	fmt.Println(x)
}

func incr(num *int) {
	*num += 1
}
