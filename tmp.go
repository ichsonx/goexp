/**
 * @Author: sonic
 * @File:  tmp.go
 * @Date: 2020/9/22 15:26
 * @Description:
 */
package main

import (
	"fmt"
)

var target = 770
var m = make(map[string]int)

// var p = []int{731, 121, 520, 214, 145, 336, 327, 372, 411, 613}
//var p = []int{731, 121, 520, 214, 145, 336, 327, 613, 260}
// var p = []int{731, 121, 214, 145, 336, 327, 613}
var p = []int{731, 214, 145, 613}

func main() {
	arr := mysubsets(p)

	r := make([]int, 0)
	for _, v := range arr {
		s := sumv(v)
		s2 := sumv(r)
		//如果结果集中的组合值的和小于target，直接换成下一个组合的
		if s2 < target {
			r = v
			continue
		}
		//如果新组合值大于target，则比较新、旧组合减去target后的值，越小的越符合要求
		//把符合的组合更新到r中
		if s >= target {
			s = s - target
			s2 = s2 - target
			if s < s2 {
				r = v
			}
		}
	}
	fmt.Printf("最佳组合：%v \n", r)
}

// 用二维数组（result）存储所有的组合，每一个元素（数组[]int）就是一个组合
// 思路：遍历二维数组，对每一个元素（数组[]int）加上遍历nums数组的一个元素，这样nums中的元素不会出现自我重叠重复的现象。
// 并且result中的每个组合都会加上新的num后会作为新的[]int元素添加到result中
// 把遍历后并加上新的nums元素的新元素（数组[]int），再添加到二维数组中作为新组合。
func mysubsets(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})
	// 每一个nums元素都往result里的组合都添加一遍
	for _, n := range nums {
		for _, arr := range result {
			arr = append(arr, n)
			result = append(result, arr)
		}
	}
	fmt.Printf("总共组合：%d 个\n", len(result))
	return result
}

//计算slice中各项之和
func sumv(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
