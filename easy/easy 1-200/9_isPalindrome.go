/**
* @Author: Huang Zengrui
* @Email: huangzengrui@yahoo.com
* @Date: 2021/8/2 19:45
* @Desc:
 */

package main

import "fmt"

func isPalindrome(x int) bool {

	origin := x
	if x < 0 {
		return false
	}
	if x/10 == 0 {
		return true
	}
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	return y == origin
}

func main() {
	fmt.Println(isPalindrome(1002))
}
