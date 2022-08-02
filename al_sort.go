package main

//冒泡排序
//比较相邻的元素。如果第一个比第二个大，就交换他们两个。
func Sort_Bubble(items []int32) {
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i] > items[j] {
				Xor_ExchangeItem(items, uint32(i), uint32(j))
			}
		}
	}
}

//首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。
//再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
//重复第二步，直到所有元素均排序完毕。
func Sort_Selected(items []int32) {
	for i := 0; i < len(items); i++ {
		tempIndex := i
		for j := i + 1; j < len(items); j++ {
			if items[tempIndex] > items[j] {
				tempIndex = j
			}
		}
		if tempIndex != i {
			Xor_ExchangeItem(items, uint32(i), uint32(tempIndex))
		}
	}
}

//插入排序
//从第1数开始，向左挨个比较，如果小，交换，如果大，停止。开始第2个数...
//有序数据，插入排序效果更好
func Sort_Insert(items []int32) {
	for i := 0; i < len(items); i++ {
		for j := i - 1; j >= 0 && items[j] > items[j+1]; j-- {
			Xor_ExchangeItem(items, uint32(j), uint32(j+1))
		}
	}
}

//归并排序
//思想：左侧排序，右侧排序，再整合
//时间复杂度：O(N*lgN)
//空间复杂度：O(N)
func Sort_Merge(items []int32) {
	if len(items) < 2 {
		return
	}
	sort_Merge_Process(items, 0, len(items)-1)
}

func sort_Merge_Process(items []int32, l, r int) {
	if l == r {
		return
	}
	mid := l + ((r - l) >> 1)
	sort_Merge_Process(items, l, mid)
	sort_Merge_Process(items, mid+1, r)
	merge(items, l, mid, r)
}

func merge(items []int32, l, m, r int) {
	temp := make([]int32, r-l+1)
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
func SmallSum(items []int32) int {
	if len(items) <= 1 {
		return 0
	}
	return smallSumProcess(items, 0, len(items)-1)
}
func smallSumProcess(items []int32, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + ((r - l) >> 1)
	return smallSumProcess(items, l, mid) + smallSumProcess(items, mid+1, r) + smallSumProcessMerge(items, l, mid, r)
}

func smallSumProcessMerge(items []int32, l, mid, r int) int {
	temp := make([]int32, r-l+1)
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
