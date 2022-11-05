package main

import "fmt"

func main() {
	maxNum := 100
	fizzNum := 3
	buzzNum := 5
	for num := 1; num < maxNum; num++ {
		if num%fizzNum == 0 && num%buzzNum == 0 {
			fmt.Println("FizzBuzz")
		} else if num%fizzNum == 0 {
			fmt.Println("Fizz")
		} else if num%buzzNum == 0 {
			fmt.Println("Buzz")
		}
	}
}
