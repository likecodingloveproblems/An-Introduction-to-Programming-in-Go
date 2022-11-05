package main

import "fmt"

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(int1 *int, int2 *int) {
	*int1, *int2 = *int2, *int1
}
