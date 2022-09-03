//暴力递归
//把大规模的问题，换成小规模的尝试
package main

import (
	"strconv"
	"strings"
)

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

//【一个整型纸牌数组，只能从左或从右抽取纸牌】，两个人轮流抽取，谁拿到的纸牌的总数多，谁获胜
//例如[1,10,100,2] ,A 1; B 10; A 100; B 1,A=101,B=11,A获胜
//返回谁获胜
func LRArrayGame(arr []int) bool {
	firstResult := lRArrayGame_first(arr, 0, len(arr)-1)
	secondResult := lRArrayGame_second(arr, 0, len(arr)-1)
	return firstResult > secondResult
}

//先拿纸牌
func lRArrayGame_first(arr []int, l, r int) int {
	//l == r 表示只有一张牌了，直接拿走
	if l == r {
		return arr[l]
	}
	//先拿左边
	resultL := arr[l] + lRArrayGame_second(arr, l+1, r)
	resultR := arr[r] + lRArrayGame_second(arr, l, r-1)
	if resultL > resultR {
		return resultL
	}
	return resultR
}
func lRArrayGame_second(arr []int, l, r int) int {
	//l == r 表示只有一张牌了，后手是0
	if l == r {
		return 0
	}
	//左边先被拿走，
	resultL := lRArrayGame_first(arr, l+1, r)
	//右边先被拿走
	resultR := lRArrayGame_first(arr, l, r-1)
	if resultL < resultR {
		return resultL
	}
	return resultR
}

//【逆序一个栈，不申请额外空间，使用递归完成】
func ReverseStack(stack *Stack) {
	if stack.isEmpty() {
		return
	}
	temp := reverseStackProcess(stack)
	ReverseStack(stack)
	stack.push(temp)
}

//将栈底元素取出
func reverseStackProcess(stack *Stack) interface{} {
	temp, _ := stack.pop()
	//如果栈是空，则temp是栈低元素，返回
	if stack.isEmpty() {
		return temp
	}
	bottom := reverseStackProcess(stack)
	stack.push(temp)
	return bottom
}

//[字符串转换成字母]，1 - A，2-B，26-Z
//例如111 可以转换成AAA，KA ，AK三种
func CalcIntString2LetterString(input string) int {
	return calcIntString2LetterStringProcess(input)
}

func calcIntString2LetterStringProcess(leftString string) int {
	if leftString == "" {
		return 1
	}
	if leftString[0:1] == "0" {
		return 0
	}
	result := 0
	if len(leftString) > 0 {
		result += calcIntString2LetterStringProcess(leftString[1:])
	}
	if len(leftString) > 1 && intString2LetterString(leftString[0:2]) != "" {
		result += calcIntString2LetterStringProcess(leftString[2:])
	}
	return result
}
func intString2LetterString(input string) string {
	inputInt, _ := strconv.ParseInt(input, 10, 64)
	if inputInt == 0 {
		return ""
	}
	if inputInt > 26 {
		return ""
	}
	return strings.ToUpper(string(rune(96 + 1)))
}
