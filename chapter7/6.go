package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(`deferred function`)
		str := recover()
		fmt.Println(str)
	}()
	fmt.Println("this is the normal statement")
	panic("PANIC")
}
