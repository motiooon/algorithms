package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	helper(nums, []int{}, &result)
	return result
}

func helper(nums []int, subres []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, subres)
		return
	}

	for i, num := range nums {
		numsCopy := append([]int{}, nums...)
		subresCopy := append([]int{}, subres...)
		helper(append(numsCopy[:i], numsCopy[i+1:]...), append(subresCopy, num), result)
	}
}

func main() {
	input := []int{1, 2, 3}
	fmt.Println(permute(input))
}