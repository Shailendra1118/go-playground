package main

import (
	"fmt"
	"testing"
)

func TestDivision(t *testing.T) {
	a := -132
	b := 6
	fmt.Println("a/b is", a/b)
	fmt.Println("b/a is", b/a)
}

func TestEvaluate(t *testing.T) {
	exp := []string{"5", "6", "+", "2", "-", "3", "/", "10", "*"}
	res := evalRPN(exp)
	fmt.Println("Result:", res)
}