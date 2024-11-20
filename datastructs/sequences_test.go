package main

import (
	"fmt"
	"math"
	"testing"
)


func findSequence(nums []int) int {
	max := 0
	boundaryMap := map[int]int{}

	for i:= 0; i<len(nums); i++ {
		if _, ok := boundaryMap[nums[i]]; !ok {
			left := boundaryMap[nums[i]-1]	
			right := boundaryMap[nums[i]+1]
			sum := left + right + 1
			boundaryMap[nums[i]] = sum
			fmt.Println("sum and boundarMap are", left, right, sum, boundaryMap)

			max = int(math.Max(float64(max), float64(sum)))

			// extend the length to the boundary(s) of the sequence
        	// will do nothing if n has no neighbors
			boundaryMap[nums[i]-left] = sum
			boundaryMap[nums[i]+right] = sum
			fmt.Println("after extending boundarMap are", boundaryMap)
		}
	}

	return max

}


func TestLargestConsecutiveSeq(t *testing.T) {
	input := []int{100, 4, 200, 1, 3, 2}
	result := findSequence(input)
	fmt.Println("Result is: ", result)
}