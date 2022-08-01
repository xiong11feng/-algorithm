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
