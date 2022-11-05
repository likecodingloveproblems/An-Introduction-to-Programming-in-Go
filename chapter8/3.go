package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

func main() {
	x := 1.5
	square(&x)
	fmt.Println(x)
	if x != 1.5*1.5 {
		panic("Something go wrong")
	}
}
