package main

import (
	"fmt"
	"math"
)

func Reverse(x int) int {
	result := 0
	for x != 0 {
		if result < math.MinInt32/10 || result > math.MaxInt32/10 {
			return 0
		}
		result = result*10 + x%10
		x = x / 10
	}
	return result
}

func main() {
	fmt.Println(Reverse(-123))
}
