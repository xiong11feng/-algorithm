//递归和非递归遍历二叉树，先序，中序，后序
//二叉树头节点是根节点
//没有孩子的节点，是叶节点
package main

import (
	"errors"
)

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

//递归遍历
//先序遍历
func PreOrderRecur(head *BinaryTree, result *[]int) {
	if head == nil {
		return
	}
	*result = append(*result, head.Value)
	PreOrderRecur(head.Left, result)
	PreOrderRecur(head.Right, result)
}

//递归遍历
//中序遍历
func InOrderRecur(head *BinaryTree, result *[]int) {
	if head == nil {
		return
	}
	InOrderRecur(head.Left, result)
	*result = append(*result, head.Value)
	InOrderRecur(head.Right, result)
}

//递归遍历
//后序遍历
func PosOrderRecur(head *BinaryTree, result *[]int) {
	if head == nil {
		return
	}
	PosOrderRecur(head.Left, result)
	PosOrderRecur(head.Right, result)
	*result = append(*result, head.Value)
}

//非递归
//思想：前序遍历的递归定义：先根节点，后左子树，再右子树。
//首先，我们遍历左子树，边遍历边打印，并把根节点存入栈中，以后需借助这些节点进入右子树开启新一轮的循环。还得重复一句：所有的节点都可看做是根节点。
func PreOrderUnRecur(head *BinaryTree, result *[]int) {
	if head == nil {
		return
	}
	stack := make(Stack, 0)
	temp := head
	for !stack.isEmpty() || temp != nil {
		for temp != nil {
			*result = append(*result, temp.Value)
			stack.push(temp)
			temp = temp.Left
		}

		for !stack.isEmpty() {
			item, _ := stack.pop()
			temp = item.(*BinaryTree)
			if temp.Right != nil {
				temp = temp.Right
				break
			}
			temp = nil
		}
	}
}

//中序遍历
//非递归思想：找到最左的节点，然后找到根，再找右侧最左节点
func InOrderUnRecur(head *BinaryTree, result *[]int) {
	if head == nil {
		return
	}
	stack := make(Stack, 0)
	temp := head
	for !stack.isEmpty() || temp != nil {
		for temp != nil {
			stack.push(temp)
			temp = temp.Left
		}
		for !stack.isEmpty() {
			item, _ := stack.pop()
			temp = item.(*BinaryTree)
			*result = append(*result, temp.Value)
			if temp.Right != nil {
				temp = temp.Right
				break
			}
			temp = nil
		}
	}
}

//后序遍历递归定义：先左子树，后右子树，再根节点。
//后序遍历的难点在于：需要判断上次访问的节点是位于左子树，还是右子树。若是位于左子树，则需跳过根节点，先进入右子树，再回头访问根节点；若是位于右子树，则直接访问根节点。
func PosOrderUnRecur(head *BinaryTree, result *[]int) {
	if head == nil {
		return
	}
	stack := make(Stack, 0)
	stackTemp := make(Stack, 0)
	temp := head
	for !stack.isEmpty() || temp != nil {
		for temp != nil {
			stack.push(temp)
			temp = temp.Left
		}
		for !stack.isEmpty() {
			//取出栈顶最左节点
			item, _ := stack.peek()
			temp = item.(*BinaryTree)

			//如果有右节点,两种情况
			if temp.Right != nil {
				tempStackItem, _ := stackTemp.peek()
				//如果此节点在临时栈中，说明，这个头节点已经向右遍历过了，无需向右子树遍历了，记录即可
				if tempStackItem != nil {
					if tempStackItem.(*BinaryTree) == temp {
						*result = append(*result, temp.Value)
						stackTemp.pop()
						stack.pop()
						temp = nil
						break
					}
				}
				//情况2，右子树没有被遍历过，当前节点入临时栈，向右子树遍历
				stackTemp.push(temp)
				temp = temp.Right
				break
			}
			//如果没有右节点，记录
			if temp.Right == nil {
				*result = append(*result, temp.Value)
				stack.pop()
			}
			temp = nil
		}
	}

}

type Stack []interface{}

// 入栈
func (s *Stack) push(a interface{}) {
	*s = append(*s, a)
}

// 出栈
func (s *Stack) pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("Empty Stack")
	}
	a := *s
	defer func() {
		*s = a[:len(a)-1]
	}()
	return a[len(a)-1], nil
}

//获取栈顶元素
func (s *Stack) peek() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("Empty Stack")
	}
	a := *s
	return a[len(a)-1], nil
}
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}
