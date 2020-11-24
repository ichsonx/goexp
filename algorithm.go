// leetcode做算法用的临时文件
package main

import (
	"fmt"
)

var lst = []int{1, 3, 5, 6}

func main() {
	fmt.Println(searchInsert(lst, 0))
}

func searchInsert(nums []int, target int) int {
	index := 0
	for i, v := range nums {
		if v > target && i == 0 {
			return 0
		}
		if v == target {
			index = i
			return i
		}
		if v < target {
			index = i
		}
	}
	return index + 1
}
