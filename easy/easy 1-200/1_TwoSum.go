package main

import "fmt"

func TwoSum(nums []int, target int) []int {

	dict := make(map[int]int)

	for idx, num := range nums {
		if value, ok := dict[target-num]; ok {
			return []int{idx, value}
		} else {
			dict[num] = idx
		}
	}
	return nil
}

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3, 4, 5, 6}, 10))
}
