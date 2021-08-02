/**
* @Author: Huang Zengrui
* @Email: huangzengrui@yahoo.com
* @Date: 2021/8/2 22:55
* @Desc:
 */

package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	lowest := math.MaxInt32

	for _, s := range strs {
		if len(s) < lowest {
			lowest = len(s)
		}
	}

	for i := 0; i < lowest; i++ {
		c := strs[0][i]
		for _, s := range strs {
			if s[i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:lowest]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flow", "flowers"}))
}
