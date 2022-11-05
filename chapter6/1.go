package main

import "fmt"

func main() {
	l := make([]int, 3, 6)
	l = append(l, 1)
	l = append(l, 2)
	l = append(l, 3)
	l = append(l, 4)
	l = append(l, 5)
	l[0] = -1
	fmt.Println(l)
}
