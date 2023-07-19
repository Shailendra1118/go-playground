package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

// Pointer receivers
// Go does not have classes, we define methods on types
// Method is a function with special 'receiver' argument
func (r *Vertex) DoubleIt() {
	r.X *= 2
	r.Y *= 2
}

func main() {
	fmt.Println("Hello Shailendra Singh Yadav")

	//closure testing
	f := fibbonacci()

	fmt.Println(f())
	fmt.Println(f())

	//struct
	v := Vertex{}
	fmt.Println(v)

	vPtr := &Vertex{}
	vPtr.X = 100
	vPtr.Y = 200
	fmt.Println(*vPtr)

	var v1 = &Vertex{}
	v1.X = 44
	v1.Y = 55
	fmt.Println("V1 pointer", v1)
	fmt.Println("V1 value", *v1)

	var v2 = &Vertex{}
	v2.X = 10
	v2.Y = 20
	v2.DoubleIt()
	fmt.Println("After doubling it", *v2)

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
