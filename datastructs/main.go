package main

import "fmt"

func init(){
	fmt.Println("Setting up some configurations...")
}

const (
	X string = "Hello"
	Y  string = "World"
	Z int = 1000;
)
func main() {
	fmt.Println("Called main() function....")
	fmt.Printf("Constants are %v, %v, %v\n", X, Y, Z)

	ch := make(chan int)
	done := make(chan bool)

	go func() {
		sum := 0
		for i := 0; i < 3; i++ {
			sum += <-ch // Receive data from channel
		}
		fmt.Println("Final Count:", sum)
		done <- false // Notify main that computation is complete
	}()

	ch <- 1 // Send data to channel
	ch <- 1 // Send data to channel
	ch <- 2 // Send data to channel

	v := <-done // Wait for the result
	fmt.Println("Got it", v)

}