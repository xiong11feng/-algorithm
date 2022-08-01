/*
递归的相关算法
*/
package main

var count = 0

//获取数组中的最大值
func GetMaxValue(items []int32) int32 {
	return getMaxValueProcess(items, 0, len(items)-1)
}

func getMaxValueProcess(items []int32, l, r int) int32 {
	count += 1
	if l == r {
		return items[l]
	}
	//取l，r中间的值
	mid := l + ((r - l) >> 1)
	lMax := getMaxValueProcess(items, l, mid)
	rMax := getMaxValueProcess(items, mid+1, r)
	if lMax > rMax {
		return lMax
	}
	return rMax
}
