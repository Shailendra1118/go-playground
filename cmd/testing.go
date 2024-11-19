package main

import "fmt"

var bar = "sdf"
const (
	item1 = iota * 20
	item2 = iota * 30
	item3
	item4 = iota
	item5

)

func main() {

	testRes := test(5,10)
	fmt.Println(testRes)
	fvar := 4.2
	res := fvar + 5
	fmt.Println("res: ", res)

	println(item1, item2, item3, item4, item5)
	fmt.Println("Hello")
	var foo = 5.0
	fmt.Println(int64(foo))

	fmt.Println("bar is:", bar)
	bar = "taP"
	fmt.Println("bar is:", bar)

	arr := []byte(bar)
	fmt.Println(arr)
	for i:=0; i <=2; i++ {
		fmt.Println("i is", i)
		if i < 100 {
			fmt.Println("i:", i, "val:", string(arr[i]))
		}
	}
}

func test(x interface{}, y interface{}) interface{} {
	switch x.(type) {
	case int:
		return y.(int) + 5
	default:
		return y
	}
}