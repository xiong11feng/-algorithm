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

func PrintTree() {

}

//二叉树的深度优先遍历就是中序遍历
//二叉树的宽度优先遍历：
//使用队列，从头进队列，弹出后，先放左孩子，再放右孩子，然后弹出
func WeigthOrder(head *BinaryTree, result *[]int) {
	if head == nil {
		return
	}
	queue := make(Queue, 0)
	temp := head
	queue.push(temp)
	for !queue.isEmpty() {
		tempItem, _ := queue.pop()
		temp = tempItem.(*BinaryTree)
		*result = append(*result, temp.Value)
		if temp.Left != nil {
			queue.push(temp.Left)
		}
		if temp.Right != nil {
			queue.push(temp.Right)
		}
	}
}

//【题目】求二叉树的最大宽度
//思路1：使用宽度优先遍历，记录每个节点所在的层数，每层统计总节点数，取得最大值
//思路2【不使用hash表】:
func MaxWeightBinaryTree(head *BinaryTree) int {
	res := 0
	if head == nil {
		return res
	}
	queue := make(Queue, 0)
	temp := head
	queue.push(temp)
	maps := make(map[*BinaryTree]int, 0)
	maps[temp] = 1
	currentLevel := 1
	currentNodes := 0
	for !queue.isEmpty() {
		tempItem, _ := queue.pop()
		temp = tempItem.(*BinaryTree)
		if currentLevel == maps[temp] {
			currentNodes++
		} else {
			if currentNodes > res {
				res = currentNodes
			}
			currentLevel = maps[temp]
			currentNodes = 1
		}
		if temp.Left != nil {
			queue.push(temp.Left)
			maps[temp.Left] = currentLevel + 1
		}
		if temp.Right != nil {
			queue.push(temp.Right)
			maps[temp.Right] = currentLevel + 1
		}
	}
	return res
}

//搜索二叉树：每颗子树，左比它小，右比它大
//【题目】如何判断一棵树是不是搜索二叉树
//思想：中序遍历，一定是升序，
//思路2:递归，左子树是否是搜索二叉树，右子树是否是搜索二叉树，左侧最大值 < 当前值 < 右侧最大值
func IsBST(head *BinaryTree) bool {
	if head == nil {
		return true
	}
	var maxValue = INT_MIN

	return isBSTRecur(head, &maxValue)
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func isBSTRecur(head *BinaryTree, maxValue *int) bool {
	if head == nil {
		return true
	}
	isleft := isBSTRecur(head.Left, maxValue)
	if !isleft {
		return false
	}
	//比较（之前中序打印的时机）
	if *maxValue > head.Value {
		return false
	}
	*maxValue = head.Value
	return isBSTRecur(head.Right, maxValue)

}

//【题目】：如何判断二叉树是完全二叉树
//宽度遍历，遇到一下情况，不是完全二叉树
//1.任意节点只有右孩子，没有左孩子
//2.如果遇到第一个左右还在不双全的情况，那么接下来的所有节点都必须是叶子节点
func IsCBT(head *BinaryTree) bool {
	if head == nil {
		return true
	}
	queue := make(Queue, 0)
	temp := head
	queue.push(temp)
	allLeaves := false
	for !queue.isEmpty() {
		tempItem, _ := queue.pop()
		temp = tempItem.(*BinaryTree)
		if allLeaves && (temp.Left != nil || temp.Right != nil) {
			return true
		}
		//1.任意节点只有右孩子，没有左孩子
		if temp.Left == nil && temp.Right != nil {
			return false
		}
		//2.如果遇到第一个左右还在不双全的情况，那么接下来的所有节点都必须是叶子节点
		if !(temp.Left != nil && temp.Right != nil) {
			allLeaves = true
		}
		if temp.Left != nil {
			queue.push(temp.Left)
		}
		if temp.Right != nil {
			queue.push(temp.Right)
		}
	}
	return true
}

//【题目】：如何判断树是满二叉树
//思路1:找出最大深度，和节点个数，节点个数 l = 2的n次方-1
//思路2:递归，左边是满二叉树，右边也是满二叉树
func IsFullBT(head *BinaryTree) bool {
	height, nodes := isFullBTProcess(head)
	return nodes == 1<<height-1
}

func isFullBTProcess(head *BinaryTree) (height, nodes int) {

	if head == nil {
		return height, nodes
	}
	height = 1
	nodes = 1
	heightMax := 0

	if head.Left != nil {
		heightL, nodesL := isFullBTProcess(head.Left)
		if heightL > heightMax {
			heightMax = heightL
		}
		nodes += nodesL
	}
	if head.Right != nil {
		heightR, nodesR := isFullBTProcess(head.Right)
		if heightR > heightMax {
			heightMax = heightR
		}
		nodes += nodesR
	}
	height = heightMax + 1
	return height, nodes
}

//【题目】：如何判断一棵树是平衡二叉树（对于任何一棵子树，它的左子树和右子树的高度差不能超过1）
func IsBalanceBT(head *BinaryTree) bool {
	result, _ := isBalanceBTPrcoess(head)
	return result
}

//树形DP题目，递归左右子树，找到解

//【题目】找到node1和node2的最低公共祖先节点
//思路1: 找到所有的节点的 parent，找node1的parent，放到集合中，
//再找node2的parent，直到出现在集合中，即为最低公共祖先节点
//
//思路2：向下递归，如果递归到的节点是空，返回空是node1，返回node1，是node2返回node2，即找到node1或者node2就提前停止递归。
//如果左右子树返回的都不空，说明返回的是node1和node2，那么当前节点就是最初的node1和node2的最低公共祖先节点，返回当前节点
//如果左右子树谁不空，就返回谁，因为不空的节点一定是node1或者node2
//换个说法，当前节点只会有三种情况：
//1.左右一个空，一个node1或者node2，是下边的情况1
//2.左右都不空，说明是node1和node2的汇集点
//3.都空，说明当前节点及其子树没有node1或者node2
//情况1: node1和node2 的最低公共祖先节点是node1或者node2
//情况2: node1和node2 的最低公共祖先节点不是node1或者node2
func LowestCommonAncester(head, node1, node2 *BinaryTree) *BinaryTree {
	if head == nil || head == node1 || head == node2 {
		//遇到nil，返回nil，
		//遇到node1，返回node1
		//遇到node2，返回node2
		return head
	}
	left := LowestCommonAncester(head.Left, node1, node2)
	right := LowestCommonAncester(head.Right, node1, node2)
	if left != nil && right != nil {
		//左右均不不是nil，返回头节点
		return head
	} else if left != nil {
		//左右两棵树并不都有返回值，left不是nil返回left，否则返回right
		return left
	} else {
		return right
	}
}

//【题目】找到一个节点的后继节点（后继节点定义，中序遍历排序后，某节点的下一个节点）
//方法一：中序遍历即可
//方法二（二叉树有parent指针）：分情况
//情况1:x有右树，后继节点是右树最左节点
//情况2:x无右树，层层向上找父节点，当当前节点是父节点的左子树的时候，此父节点就是后继节点
//（原因是，x是此节点的左子树的最右节点，x后一定是此节点），找不到的话，后继节点就是空

func isBalanceBTPrcoess(head *BinaryTree) (isBanlance bool, height int) {
	if head == nil {
		return true, 0
	}
	height = 1
	heightLeft := 0
	heightRight := 0
	if head.Left != nil {
		isBanlance, heightLeft = isBalanceBTPrcoess(head.Left)
		if !isBanlance {
			return false, height
		}
	}
	if head.Right != nil {
		isBanlance, heightRight = isBalanceBTPrcoess(head.Right)
		if !isBanlance {
			return false, height
		}
	}
	if heightLeft-heightRight > 1 || heightLeft-heightRight < -1 {
		return false, height
	}
	if heightRight > heightLeft {
		height += heightRight
	} else {
		height += heightLeft
	}
	return true, height
}

//【题目】打印纸条对摺n次后的凹凸折痕
//n=1，凹折痕
//n=2，凹凹凸折痕
//n=3 ...
//折纸发现规律，n其实是二叉树的高度，头节点是凹，每一个左子树的头节点都是凹，每一个右子树的头节点都是凸
//纸条上的折痕顺序，其实就是上述二叉树的中序遍历的结果
//解决方案：
//1.中序遍历二叉树，打印，额外空间较多
//2.递归的方式

func FoldPaper(n int) []int {
	result := make([]int, 0)
	folderPaperProcess(n, 1, true, &result)
	return result
}

func folderPaperProcess(n, height int, down bool, arr *[]int) {
	if n < height {
		return
	}
	//左侧递归
	folderPaperProcess(n, height+1, true, arr)
	//当前值进数组
	if down {
		*arr = append(*arr, 1)
	} else {
		*arr = append(*arr, 0)
	}
	folderPaperProcess(n, height+1, false, arr)
}

type Queue []interface{}

func (q *Queue) push(a interface{}) {
	*q = append(*q, a)
}
func (q *Queue) pop() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("Empty Queue")
	}
	a := *q
	defer func() {
		*q = a[1:]
	}()
	return a[0], nil
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
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
