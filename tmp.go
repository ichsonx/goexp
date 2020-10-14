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
//var p = []int{731, 121, 520, 214, 145, 336, 327, 372, 411, 613}
var p = []int{731, 121, 520, 214, 145, 336, 327, 613}

func main() {
	arr := subsets(p)

	r := make([]int,0)
	for _, v := range arr{
		s := sumv(v)
		s2 := sumv(r)
		//如果结果集中的组合值的和小于target，直接换成下一个组合的
		if s2 < target{
			r = v
			continue
		}
		//如果新组合值大于target，则比较新、旧组合减去target后的值，越小的越符合要求
		//把符合的组合更新到r中
		if s >= target  {
			s = s - target
			s2 = s2 - target
			if s < s2 {
				r = v
			}
		}
	}
	fmt.Println(arr)
	fmt.Println(r)
}

func subsets(nums []int) [][]int {
	arr:=make([][]int,0)
	arr=append(arr,[]int{})
	for i:=0;i<len(nums);i++{
		tempArr:=make([][]int,0)
		for _,c:=range arr{
			temp:=make([]int,0)
			temp=append(temp,c...)
			temp=append(temp,nums[i])
			fmt.Println(temp)
			tempArr=append(tempArr,temp)
		}
		for _,c:=range tempArr{
			arr=append(arr,c)
		}
	}
	return arr[1:]
}

//计算slice中各项之和
func sumv(a []int) int{
	sum := 0
	for _,v := range a{
		sum += v
	}
	return sum
}
