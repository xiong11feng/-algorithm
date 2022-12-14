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

// integerHeap a type
type IntegerHeap []int

// IntegerHeap method - gets the length of integerHeap
func (iheap IntegerHeap) Len() int { return len(iheap) }

// IntegerHeap method - checks if element of i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

// IntegerHeap method -swaps the element of i to j index
func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

//IntegerHeap method -pushes the item
func (iheap *IntegerHeap) Push(heapintf interface{}) {

	*iheap = append(*iheap, heapintf.(int))
}

//IntegerHeap method -pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}
