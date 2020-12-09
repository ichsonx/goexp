// twoTree
package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t TreeNode) display() {
	fmt.Printf("节点值：%d\n", t.Val)
}

func (t *TreeNode) AddNode(val int) error {
	if t.Left != nil && t.Right != nil {
		return errors.New(fmt.Sprintf("error：当前节点度数为2（已满），不能再添加子节点\n"))
	}

	newNode := TreeNode{Val: val}
	if t.Left == nil {
		t.Left = &newNode
	} else {
		t.Right = &newNode
	}
	return nil
}
func (t *TreeNode) AddLeftNode(val int) error {
	if t.Left != nil {
		return errors.New(fmt.Sprintf("error：左子节点不为空\n"))
	} else {
		newNode := TreeNode{Val: val}
		t.Left = &newNode
	}
	return nil
}
func (t *TreeNode) AddRightNode(val int) error {
	if t.Right != nil {
		return errors.New(fmt.Sprintf("error：右子节点不为空\n"))
	} else {
		newNode := TreeNode{Val: val}
		t.Right = &newNode
	}
	return nil
}

func (t TreeNode) duDisplay() {
	du := 0
	if t.Right != nil {
		du = 2
	} else if t.Left != nil {
		du = 1
	}
	fmt.Printf("度数：%d\n", du)
}

func main() {
	fmt.Println("Hello World!")
	t := TreeNode{Val: 1}
	t.AddNode(5)
	err := t.AddNode(9)
	if err != nil {
		fmt.Print(err)
	}
	t.display()
	t.duDisplay()
}
