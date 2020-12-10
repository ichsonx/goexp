// twoTree
package main

import (
	"errors"
	"fmt"
)

// 二叉树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 显示当前节点的值
func (t TreeNode) display() {
	fmt.Printf("节点值：%d\n", t.Val)
}

// 自动添加节点，顺序从左到右。返回error
func (t *TreeNode) AddNode(val int) (err error) {
	if t.Left != nil && t.Right != nil {
		return errors.New(fmt.Sprintf("error：当前节点度数为2（已满），不能再添加子节点\n"))
	}

	if err = t.AddLeftNode(val); err != nil {
		err = t.AddRightNode(val)
	}

	return err
}

// 添加左节点，返回error
func (t *TreeNode) AddLeftNode(val int) error {
	if t.Left != nil {
		return errors.New(fmt.Sprintf("error：左子节点不为空\n"))
	} else {
		newNode := TreeNode{Val: val}
		t.Left = &newNode
	}
	return nil
}

// 添加右节点，返回error
func (t *TreeNode) AddRightNode(val int) error {
	if t.Right != nil {
		return errors.New(fmt.Sprintf("error：右子节点不为空\n"))
	} else {
		newNode := TreeNode{Val: val}
		t.Right = &newNode
	}
	return nil
}

// 返回当前节点的度
func (t TreeNode) Du() int {
	du := 0
	if t.Right != nil {
		du = 2
	} else if t.Left != nil {
		du = 1
	}
	return du
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
}
