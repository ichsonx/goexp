// twoTree
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t TreeNode) display() {
	fmt.Printf("节点值：%d/n", t.Val)
}

func main() {
	fmt.Println("Hello World!")
}
