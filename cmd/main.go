package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Shailendra Singh Yadav")

	//closure testing
	f := fibbonacci()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func fibbonacci() func() int {
	last := 1
	secondLast := 0

	return func() int {
		current := last + secondLast
		secondLast = last
		last = current
		return last
	}
}
