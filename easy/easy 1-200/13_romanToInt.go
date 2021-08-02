/**
* @Author: Huang Zengrui
* @Email: huangzengrui@yahoo.com
* @Date: 2021/8/2 22:00
* @Desc:
 */

package main

import "fmt"

var relationMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := relationMap[s[i]]
		if i < n-1 && value < relationMap[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}
func main() {
	fmt.Println(romanToInt("XIV"))
}
