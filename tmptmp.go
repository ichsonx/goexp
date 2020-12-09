// tmptmp
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	z := [][]int{}
	z = append(z, []int{})

	for _, n := range nums {
		for _, a := range z {
			a = append(a, n)
			z = append(z, a)
		}
	}
	fmt.Println(z)
	fmt.Println(len(z))
}
