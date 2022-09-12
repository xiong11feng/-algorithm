package main

//【机器人走步问题】
// 从1～N的格子，给定开始个字start，结束个字end，（1<= start <=N  1<= end <=N）
// 机器人每次走1格, 必须走k步，左走右走均可以，到1之后，只能走2，N之后只能走N-1
// 求一共有多少种走法

//暴力递归的解法
func RobotWalk1(k, n, start, end int) int {
	return robotWalk1Process(n, end, k, start)
}
func robotWalk1Process(n, end, remainStep, curPos int) int {
	if remainStep == 0 {
		if curPos == end {
			return 1
		} else {
			return 0
		}
	}
	if curPos == 1 {
		return robotWalk1Process(n, end, remainStep-1, 2)
	}
	if curPos == n {
		return robotWalk1Process(n, end, remainStep-1, n-1)
	}
	return robotWalk1Process(n, end, remainStep-1, curPos+1) +
		robotWalk1Process(n, end, remainStep-1, curPos-1)
}

func RobotWalk2(k, n, start, end int) int {
	//缓存,某点的某步数，值是相同的，无需重复计算
	//k*n的矩阵
	cache := make([][]int, k+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	for i := 0; i < k+1; i++ {
		for j := 0; j < n+1; j++ {
			cache[i][j] = -1 //-1表示没有缓存
		}
	}
	return robotWalk2Process(n, end, k, start, &cache)
}

func robotWalk2Process(n, end, remainStep, curPos int, cache *[][]int) int {
	if (*cache)[remainStep][curPos] != -1 {
		return (*cache)[remainStep][curPos]
	}
	if remainStep == 0 {
		if curPos == end {
			(*cache)[remainStep][curPos] = 1
			return 1
		} else {
			(*cache)[remainStep][curPos] = 0
			return 0
		}
	}
	temp := 0
	if curPos == 1 {
		temp = robotWalk2Process(n, end, remainStep-1, 2, cache)
	} else if curPos == n {
		temp = robotWalk2Process(n, end, remainStep-1, n-1, cache)
	} else {
		temp = robotWalk2Process(n, end, remainStep-1, curPos+1, cache) +
			robotWalk2Process(n, end, remainStep-1, curPos-1, cache)
	}
	(*cache)[remainStep][curPos] = temp
	return temp

}
