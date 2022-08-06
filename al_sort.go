package main

import (
	"container/heap"
	"math/rand"
)

//冒泡排序
//比较相邻的元素。如果第一个比第二个大，就交换他们两个。
func Sort_Bubble(items []int) {
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i] > items[j] {
				Xor_ExchangeItem(items, int(i), int(j))
			}
		}
	}
}

//首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。
//再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
//重复第二步，直到所有元素均排序完毕。
func Sort_Selected(items []int) {
	for i := 0; i < len(items); i++ {
		tempIndex := i
		for j := i + 1; j < len(items); j++ {
			if items[tempIndex] > items[j] {
				tempIndex = j
			}
		}
		if tempIndex != i {
			Xor_ExchangeItem(items, int(i), int(tempIndex))
		}
	}
}

//插入排序
//从第1数开始，向左挨个比较，如果小，交换，如果大，停止。开始第2个数...
//有序数据，插入排序效果更好
func Sort_Insert(items []int) {
	for i := 0; i < len(items); i++ {
		for j := i - 1; j >= 0 && items[j] > items[j+1]; j-- {
			Xor_ExchangeItem(items, int(j), int(j+1))
		}
	}
}

//归并排序
//思想：左侧排序，右侧排序，再整合
//时间复杂度：O(N*lgN)
//空间复杂度：O(N)
func Sort_Merge(items []int) {
	if len(items) < 2 {
		return
	}
	sort_Merge_Process(items, 0, len(items)-1)
}

func sort_Merge_Process(items []int, l, r int) {
	if l == r {
		return
	}
	mid := l + ((r - l) >> 1)
	sort_Merge_Process(items, l, mid)
	sort_Merge_Process(items, mid+1, r)
	merge(items, l, mid, r)
}

func merge(items []int, l, m, r int) {
	temp := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := m + 1
	for p1 <= m && p2 <= r {
		if items[p1] < items[p2] {
			temp[i] = items[p1]
			i++
			p1++
		} else {
			temp[i] = items[p2]
			i++
			p2++
		}
	}
	for p1 <= m {
		temp[i] = items[p1]
		i++
		p1++
	}
	for p2 <= r {
		temp[i] = items[p2]
		i++
		p2++
	}
	for i := 0; i < len(temp); i++ {
		items[l+i] = temp[i]
	}
}

//求小和问题
//在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。求一个数组的小和。
//例如[1,3,5,4] 1小和是0，3的小和1，5的小和1+3，4的小和1+3 ，所以数组的小和是9
//算法思路：左侧小和转换成右侧大于本身的个数 * 自身的值，1 右侧三个大于1的数，小和是1*3，3右侧是2*3，5右侧是0，4右侧是0，数组小和是9
//使用归并排序，归并的过程中，计算出小和
func SmallSum(items []int) int {
	if len(items) <= 1 {
		return 0
	}
	return smallSumProcess(items, 0, len(items)-1)
}
func smallSumProcess(items []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + ((r - l) >> 1)
	return smallSumProcess(items, l, mid) + smallSumProcess(items, mid+1, r) + smallSumProcessMerge(items, l, mid, r)
}

func smallSumProcessMerge(items []int, l, mid, r int) int {
	temp := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := mid + 1
	result := 0
	for p1 <= mid && p2 <= r {
		if items[p1] < items[p2] {
			temp[i] = items[p1]
			result += (r - p2 + 1) * int(temp[i])
			i++
			p1++

		} else {
			temp[i] = items[p2]
			i++
			p2++
		}
	}
	for p1 <= mid {
		temp[i] = items[p1]
		i++
		p1++
	}
	for p2 <= r {
		temp[i] = items[p2]
		i++
		p2++
	}
	for j := 0; j < len(temp); j++ {
		items[l+j] = temp[j]
	}
	return result
}

//荷兰国旗问题
//设有一个数组，需要将小于数字a的数，放在数组左边，等于a的数放在中间，大于a的放在右边
//算法思路：
//假设数组长度是n，初始l=-1（l表示小于a的区域），r=n（r表示大于a的区域），所以l～r之间是等于a的区域
//从头开始遍历数组，i=0开始，到i=n-1，如果：
//arr[i] < a, swap(arr[l+1],arr[i]),i++,l++
//arr[i] == a, i++
//arr[i] > a, swap(arr[r-1],arr[i]),r--,i不变

func DutchFlag(items []int, a int) {
	l := -1
	r := len(items)
	for i := 0; i < r; {
		if items[i] < a {
			if i != l+1 {
				Xor_ExchangeItem(items, int(i), int(l+1))
			}
			l++
			i++
		} else if items[i] == a {
			i++
		} else {
			if i != r-1 {
				Xor_ExchangeItem(items, int(i), int(r-1))
			}
			r--
		}
	}
}

//快速排序2.0版本，是荷兰国旗问题的拓展
//假设数组是[4,3,5,6,5,0,1,7,8,5]
//以数组最后一位作为a，分成三份[<5,=5,>5]
//等于5的份，已经排序完成，左右两份继续相同办法，递归排序
//最坏情况：数组已经有序，事件复杂度，O(N平方)
//选择的a值，最好的情况是a在组中排在中等，O(NlgN)
//a如果总是在左右边上，是最坏情况，O(N平方)
func Sort_Fast_V2(items []int) {
	if len(items) < 2 {
		return
	}
	sort_Fast_V2_Process(items, 0, len(items)-1)
}
func sort_Fast_V2_Process(items []int, left, right int) {
	if left == right {
		return
	}
	l := left - 1
	r := right + 1
	a := items[right]
	for i := left; i < r; {
		if items[i] < a {
			if i != l+1 {
				Xor_ExchangeItem(items, int(i), int(l+1))
			}
			l++
			i++
		} else if items[i] == a {
			i++
		} else {
			if i != r-1 {
				Xor_ExchangeItem(items, int(i), int(r-1))
			}
			r--
		}
	}
	if left < l {
		//左侧递归
		sort_Fast_V2_Process(items, left, l)
	}
	if r < right {
		//右侧递归
		sort_Fast_V2_Process(items, r, right)
	}

}

//快速排序3.0
//解决a值可能总是最差情况
//3.0不使用末尾值作为a，而是随机选择一个值，作为a
//这样好情况和坏情况变成了概率事件
//计算出概率的期望值是O(NlgN)
func Sort_Fast_V3(items []int) {
	if len(items) < 2 {
		return
	}
	sort_Fast_V3_Process(items, 0, len(items)-1)
}
func sort_Fast_V3_Process(items []int, left, right int) {
	if left == right {
		return
	}

	//随机选一个a
	aIndex := rand.Intn(right+1-left) + left
	if aIndex != right {
		Xor_ExchangeItem(items, int(right), int(aIndex))
	}

	l := left - 1
	r := right + 1
	a := items[right]
	for i := left; i < r; {
		if items[i] < a {
			if i != l+1 {
				Xor_ExchangeItem(items, int(i), int(l+1))
			}
			l++
			i++
		} else if items[i] == a {
			i++
		} else {
			if i != r-1 {
				Xor_ExchangeItem(items, int(i), int(r-1))
			}
			r--
		}
	}
	if left < l {
		//左侧递归
		sort_Fast_V2_Process(items, left, l)
	}
	if r < right {
		//右侧递归
		sort_Fast_V2_Process(items, r, right)
	}

}

//堆排序
//利用大根堆堆顶记录的是最大关键字(最小关键字)这一特性，使得每次从无序中选择最大记录(最小记录)变得简单。
//将待排序的序列构造成一个最大堆，此时序列的最大值为根节点
//依次将根节点与待排序序列的最后一个元素交换
//再维护从根节点到该元素的前一个节点为最大堆，如此往复，最终得到一个递增序列
func Sort_Heap(items []int) {
	if len(items) < 2 {
		return
	}

	//先找到大根堆
	for i := 1; i < len(items); i++ {
		HeapInsert(items, i)
	}
	Xor_ExchangeItem(items, 0, len(items)-1)
	heapSize := len(items) - 1
	for heapSize > 1 {
		Heapify(items, 0, heapSize)
		heapSize--
		Xor_ExchangeItem(items, 0, heapSize)
	}
}

//堆排序的拓展
//如果一个基本有序的数组，排序后每个元素移动距离不超过k，如何进行排序
//题目思考：使用堆排序，从坐标0开始，到k-1，先组成小根堆，之后弹出索引，0，压入k ...
//弹出1，压入k+1 直到结束

func SortArrDistanceLessK(items []int, k int) {
	h := &IntegerHeap{}
	for i := 0; i < k; i++ {
		heap.Push(h, items[i])
	}
	length := len(items)
	top := 0
	tail := k
	for top < len(items) {
		items[top] = heap.Pop(h).(int)
		if tail < length {
			heap.Push(h, items[tail])
		}
		top++
		tail++
	}
}

//桶排序 - 基数排序
//算法思想：
//1.找到数组中最大的数的位数
//2.从最后低位开始排序，一轮一轮排序，一直排到最高
func Sort_Bucket(items []int) {
	if len(items) < 2 {
		return
	}
	maxBits := getMaxDigitsBits(items)
	radix := make([]int, 10)
	help := make([]int, len(items))

	for i := 1; i <= maxBits; i++ {
		for j := 0; j < len(items); j++ {
			x := getDigits(items[j], i)
			radix[x]++
		}
		//将radix变成后一项是前边n项的和
		for j := 1; j < len(radix); j++ {
			radix[j] = radix[j-1] + radix[j]
		}
		//从后遍历原数组，根据radix，将每个数组调整位置
		for j := len(items) - 1; j >= 0; j-- {
			x := getDigits(items[j], i)
			help[radix[x]-1] = items[j]
			//减1目前，下次遇到相同的值时，放在更前一个位置，即不改变排好的顺序
			radix[x]--
		}
		//再将help数组copy回items
		for j := 0; j < count; j++ {
			items[j] = help[j]
		}
	}

}

//获取一个int的某位的值，1表示个位，2表示10位，3表示百位...
func getDigits(item int, bit int) int {
	result := item
	for result != 0 && bit > 0 {
		a := bit
		sqrt := 10
		for a > 1 {
			sqrt *= 10
			a--
		}
		result %= sqrt
		bit--
	}
	return result
}

func getMaxDigitsBits(items []int) int {
	max := items[0]
	for i := 1; i < len(items); i++ {
		if max < items[i] {
			max = items[i]
		}
	}
	j := 0
	for max > 0 {
		max = max / 10
		j++
	}
	return j
}
