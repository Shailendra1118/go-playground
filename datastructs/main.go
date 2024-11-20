package main

import (
	"fmt"
	"strconv"
)

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

	//v := <-done // Wait for the result
	//fmt.Println("Got it", v)
	fmt.Println("End of this...")

}



//----------------

func evalRPN(tokens []string) int {
    var stack []string
    //map literals
    operators := map[string]struct{}{
        "+": {},
        "-": {},
		"*": {},
		"/": {},
    }
    for i :=0; i<len(tokens); i++ {
        token := tokens[i]
        if _, ok := operators[token]; !ok {
            stack = append(stack, token)
        } else {
            first := stack[len(stack)-1]
			//stack = stack[:len(stack)-1]
            sec := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
            newToken := evaluate(token, sec, first)
            stack = append(stack, strconv.Itoa(newToken))
        }
    }

	fmt.Println("Stack has:",stack)
    res, _ := strconv.Atoi(stack[0])
	return res

}

func evaluate(operator, operand1, operand2 string) int {
	op1Int, _ := strconv.Atoi(operand1)
	op2Int, _ := strconv.Atoi(operand2)
    switch operator {
        case "+":
			return op1Int + op2Int
        case "-":
            return op1Int - op2Int
		case "*":
            return op1Int * op2Int
		case "/":
            return op1Int / op2Int
        default:
            return 0
		}
}