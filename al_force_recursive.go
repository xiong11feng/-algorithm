//暴力递归
//把大规模的问题，换成小规模的尝试
package main

//【汉诺塔问题】n层汉诺塔，打印所有轨迹
//递归思想：
//假设有n个盘，A，B，C 三个柱，启示A柱上有n个盘
//1.将A上n-1个盘，移动到B
//2.A上第n个最大的盘，移到C
//3.将b上n-1个盘，移动到C
func Hanoi(n int) int {
	return hanoi_move(n, "A", "B", "C")
}
func hanoi_move(n int, from, to, other string) int {
	result := 0
	if n > 1 {
		result += hanoi_move(n-1, from, other, to)
	}
	result += 1
	if n > 1 {
		result += hanoi_move(n-1, other, to, from)
	}
	return result
}

//【打印一个字符串的全部子序列，包括空字符串】
// 二叉树，abc，每个位置都有两条路，后边的字符要或者不要
// 子序列，可以不连续
func SubSequenceString(a string, result *[]string) {
	subSequenceStringProcess(a, "", 0, result)
}
func subSequenceStringProcess(a string, current string, i int, result *[]string) {

	if len(a) < i+1 {
		//log.Println("current string is: " + current)
		*result = append(*result, current)
		return
	}
	temp := a[i : i+1]

	current += temp
	subSequenceStringProcess(a, current, i+1, result) //要当前字符

	current = current[:len(current)-1]
	subSequenceStringProcess(a, current, i+1, result) //不要当前字符

}

//【找到字符串的全排列】
func FindAllPermutation(input string, result *[]string) {
	findAllPermutationProcess("", input, result)
}

//current表示之前已经拼接好的
//needPermutation 表示还剩下没有处理的
func findAllPermutationProcess(current string, needPermutation string, result *[]string) {
	//如果只剩下最后一个字符，无需排序了，拼接到current
	if len(needPermutation) == 1 {
		current += needPermutation
		*result = append(*result, current)
		return
	}
	//判断字符串是不是重复，
	//还没有拼接的字符串needPermutation，如果有重复项，只处理一个即可。
	//处理多次会重复
	duplicated := make(map[byte]struct{}, 0)
	tempCurrent := current
	for j := 0; j < len(needPermutation); j++ {
		if _, ok := duplicated[needPermutation[j]]; !ok {
			findAllPermutationProcess(tempCurrent+string(needPermutation[j]), needPermutation[:j]+needPermutation[j+1:], result)
		}
	}
}
