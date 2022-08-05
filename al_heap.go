//Heap 堆的相关问题
//完全二叉树：没有左分支，而右分支，就不是完全二叉树，其余都是完全二叉树
//定义：一棵深度为k的有n个结点的二叉树，对树中的结点按从上至下、从左到右的顺序进行编号，如果编号为i（1≤i≤n）的结点与满二叉树中编号为i的结点在二叉树中的位置相同，
//堆是特殊的完全二叉树
//大根堆，树（子树）的最大值，是头节点
//小根堆，树（子树）的最小值，是头节点
package main

//index位置的节点，向上移动
//二叉树，父节点的坐标是（index-1）/2
func HeapInsert(items []int, index int) {
	for items[index] > items[(index-1)/2] {
		Xor_ExchangeItem(items, index, int((index-1)/2))
		index = (index - 1) / 2
	}
}

//某个数在index位置，能否向下移动
func Heapify(items []int, index, heapSize int) {
	//左孩子坐标
	left := index*2 + 1
	for left < heapSize {
		maxChildIndex := left
		if left+1 < heapSize && items[left+1] > items[maxChildIndex] {
			maxChildIndex = left + 1
		}
		if items[index] < items[maxChildIndex] {
			Xor_ExchangeItem(items, int(index), int(maxChildIndex))
		}
		index = maxChildIndex
		left = 2*index + 1
	}
}
