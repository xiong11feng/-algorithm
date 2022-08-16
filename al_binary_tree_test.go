package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//测试树
//                 1
//             /       \
//            2         7
//          /   \         \
//         3     4         8
//                \         \
//                 5         9
//                /
//               6

func GetTestTree() *BinaryTree {
	tree := &BinaryTree{Value: 1}
	tree.Left = &BinaryTree{Value: 2}
	tree.Left.Left = &BinaryTree{Value: 3}
	tree.Left.Right = &BinaryTree{Value: 4}
	tree.Left.Right.Right = &BinaryTree{Value: 5}
	tree.Left.Right.Right.Left = &BinaryTree{Value: 6}
	tree.Right = &BinaryTree{Value: 7}
	tree.Right.Right = &BinaryTree{Value: 8}
	tree.Right.Right.Right = &BinaryTree{Value: 9}
	return tree
}

//测试树
//                 7
//             /       \
//            3         8
//          /   \         \
//         2     4         9
//                \         \
//                6          10
//                /
//               5
func GetBSTree() *BinaryTree {
	tree := &BinaryTree{Value: 7}
	tree.Left = &BinaryTree{Value: 3}
	tree.Left.Left = &BinaryTree{Value: 2}
	tree.Left.Right = &BinaryTree{Value: 4}
	tree.Left.Right.Right = &BinaryTree{Value: 6}
	tree.Left.Right.Right.Left = &BinaryTree{Value: 5}
	tree.Right = &BinaryTree{Value: 8}
	tree.Right.Right = &BinaryTree{Value: 9}
	tree.Right.Right.Right = &BinaryTree{Value: 10}
	return tree
}

//测试树
//                 7
//             /       \
//            3         8
//          /   \
//         2     4
func GetCBTree() *BinaryTree {
	tree := &BinaryTree{Value: 7}
	tree.Left = &BinaryTree{Value: 3}
	tree.Left.Left = &BinaryTree{Value: 2}
	tree.Left.Right = &BinaryTree{Value: 4}
	tree.Right = &BinaryTree{Value: 8}
	return tree
}
func TestPreOrderRecur(t *testing.T) {
	tree := GetTestTree()
	actual := make([]int, 0)
	PreOrderRecur(tree, &actual)
	assert.Equal(t, "123456789", getArrString(actual))
}

func TestPreOrderUnRecur(t *testing.T) {
	tree := GetTestTree()
	actual := make([]int, 0)
	PreOrderUnRecur(tree, &actual)
	assert.Equal(t, "123456789", getArrString(actual))
}

func TestInOrderRecur(t *testing.T) {
	tree := GetTestTree()
	actual := make([]int, 0)
	InOrderRecur(tree, &actual)
	assert.Equal(t, "324651789", getArrString(actual))
}

func TestInOrderUnRecur(t *testing.T) {
	tree := GetTestTree()
	actual := make([]int, 0)
	InOrderUnRecur(tree, &actual)
	assert.Equal(t, "324651789", getArrString(actual))
}

func TestPosOrderRecur(t *testing.T) {
	tree := GetTestTree()
	actual := make([]int, 0)
	PosOrderRecur(tree, &actual)
	assert.Equal(t, "365429871", getArrString(actual))
}

func TestPosOrderUnRecur(t *testing.T) {
	tree := GetTestTree()
	actual := make([]int, 0)
	PosOrderUnRecur(tree, &actual)
	assert.Equal(t, "365429871", getArrString(actual))
}
func TestWeigthOrder(t *testing.T) {
	tree := GetTestTree()
	actual := make([]int, 0)
	WeigthOrder(tree, &actual)
	assert.Equal(t, "127348596", getArrString(actual))
}

func TestMaxWeightBinaryTree(t *testing.T) {
	tree := GetTestTree()
	actual := MaxWeightBinaryTree(tree)
	assert.Equal(t, 3, actual)
}

func TestIsBST(t *testing.T) {
	tree := GetTestTree()
	actual := IsBST(tree)
	assert.Equal(t, false, actual)

	tree = GetBSTree()
	actual = IsBST(tree)
	assert.Equal(t, true, actual)
}

func TestIsCBT(t *testing.T) {
	tree := GetTestTree()
	actual := IsCBT(tree)
	assert.Equal(t, false, actual)

	tree = GetCBTree()
	actual = IsCBT(tree)
	assert.Equal(t, true, actual)
}

func getArrString(arr []int) string {
	res := ""
	for i := 0; i < len(arr); i++ {
		res += fmt.Sprintf("%d", arr[i])
	}
	return res
}
