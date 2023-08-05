package pointer

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	fmt.Println("modifying v in Abs")
	v.X = 100
	return v.X * v.X
}

func (v Vertex) Abs1() float64 {
	fmt.Println("Mofifying v in Abs1")
	v.X = 50
	return v.X * v.X
}

func main() {
	v := Vertex{1, 1}
	fmt.Println(v.Abs())
	fmt.Printf("After calling Abs() value of X is %v\n", v.X)

	v1 := &Vertex{2, 2}
	fmt.Println((*v1).Abs1()) // == v1.Abs1()
	fmt.Printf("After calling Abs1() value of X is %v\n", v.X)
}
