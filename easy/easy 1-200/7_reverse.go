package main

import "fmt"

func Reverse(x int) int {
	result := 0
	for ; x/10 > 0; x = x / 10 {
		result = x%10 + result*10
	}
	result = x%10 + result*10
	return result
}

func main() {
	fmt.Println(Reverse(-123))
}
