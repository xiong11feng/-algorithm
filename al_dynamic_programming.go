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

//使用记忆，防止重复计算
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

//【象棋游戏】
//9*10的棋盘，x坐标（0，8），y（0，9）
//马 从（0，0）点开始去 指定点（x,y）,必须走k步，有多少种走法

func HorseGame1(x, y, k int) int {
	return horseGameProcess1(x, y, k)
}

//相当于从（x，y）点想（0，0）点递归
func horseGameProcess1(x, y, step int) int {
	if step == 0 {
		if x == 0 && y == 0 {
			return 1
		}
		return 0
	}
	//9*10的棋盘
	if x > 8 || x < 0 || y > 9 || y < 0 {
		return 0
	}
	//马可以跳8个方向
	res := horseGameProcess1(x-1, y+2, step-1) +
		horseGameProcess1(x+1, y+2, step-1) +
		horseGameProcess1(x+2, y+1, step-1) +
		horseGameProcess1(x+2, y-1, step-1) +
		horseGameProcess1(x+1, y-2, step-1) +
		horseGameProcess1(x-1, y-2, step-1) +
		horseGameProcess1(x-2, y-1, step-1) +
		horseGameProcess1(x-2, y+1, step-1)
	return res
}

//记忆优化
func HorseGame2(x, y, k int) int {
	//三维数组，想像成立方体
	cache := make([][][]int, k+1)
	for a := 0; a < k+1; a++ {
		cache[a] = make([][]int, 9)
		for b := 0; b < 9; b++ {
			cache[a][b] = make([]int, 10)
			for c := 0; c < 10; c++ {
				cache[a][b][c] = 0
			}
		}
	}
	cache[0][0][0] = 1

	for a := 0; a < k; a++ {
		for b := 0; b < 9; b++ {
			for c := 0; c < 10; c++ {
				if cache[a][b][c] != 0 {
					temp := cache[a][b][c]
					//上层的八个点，不越界的情况下，都会受影响
					if b+2 <= 8 && c+1 <= 9 {
						cache[a+1][b+2][c+1] += temp
					}
					if b+2 <= 8 && c-1 >= 0 {
						cache[a+1][b+2][c-1] += temp
					}
					if b+1 <= 8 && c+2 <= 9 {
						cache[a+1][b+1][c+2] += temp
					}
					if b+1 <= 8 && c-2 >= 0 {
						cache[a+1][b+1][c-2] += temp
					}
					if b-1 >= 0 && c+2 <= 9 {
						cache[a+1][b-1][c+2] += temp
					}
					if b-1 >= 0 && c-2 >= 0 {
						cache[a+1][b-1][c-2] += temp
					}
					if b-2 >= 0 && c+1 <= 9 {
						cache[a+1][b-2][c+1] += temp
					}
					if b-2 >= 0 && c-1 >= 0 {
						cache[a+1][b-2][c-1] += temp
					}
				}
			}
		}

	}
	return cache[k][x][y]
}

//[破钱问题]
//有一个数组，[1,2,3,5,10,20,50,100]，每个值代表货币的面值
//目标价 aim 元，每种货币有无限张，求有多少种方法，组成aim元

func FindMoney1(aim int, arr []int) int {
	return findMoney1Process(0, aim, arr)
}

//index 数组从index开始的钱，可以用
//aim，目标钱数
//arr 钱数数组
func findMoney1Process(index int, aim int, arr []int) int {
	if index == len(arr) {
		if aim == 0 {
			return 1
		}
		return 0
	}
	res := 0
	for i := 0; i*arr[index] <= aim; i++ {
		res += findMoney1Process(index+1, aim-i*arr[index], arr)
	}
	return res
}

//记忆搜索,精确DP
func FindMoney2(aim int, arr []int) int {
	dp := make([][]int, len(arr)+1) //面值
	for i := 0; i < len(arr)+1; i++ {
		dp[i] = make([]int, aim+1)
		for j := 0; j < aim+1; j++ {
			dp[i][j] = 0
		}
	}
	dp[len(arr)][0] = 1 //index == len(arr)
	//每一个点dp[x][*] 是有其dp[x+1][*]决定的
	for index := len(arr) - 1; index >= 0; index-- {
		for j := 0; j < aim+1; j++ {
			// for i := 0; i*arr[index] <= aim; i++ {
			// 	res += findMoney1Process(index+1, aim-i*arr[index], arr)
			// }
			//意义同上述代码
			res := 0
			for i := 0; i*arr[index] <= j; i++ {
				res += dp[index+1][j-i*arr[index]]
			}
			dp[index][j] = res
		}
	}
	return dp[0][aim]
}

//继续优化搜索
func FindMoney3(aim int, arr []int) int {
	dp := make([][]int, len(arr)+1) //面值
	for i := 0; i < len(arr)+1; i++ {
		dp[i] = make([]int, aim+1)
		for j := 0; j < aim+1; j++ {
			dp[i][j] = 0
		}
	}
	dp[len(arr)][0] = 1 //index == len(arr)
	//每一个点dp[x][*] 是有其dp[x+1][*]决定的
	for index := len(arr) - 1; index >= 0; index-- {
		for j := 0; j < aim+1; j++ {
			// for i := 0; i*arr[index] <= aim; i++ {
			// 	res += findMoney1Process(index+1, aim-i*arr[index], arr)
			// }
			//意义同上述代码
			//优化，同行的点，的前arr[index]点已经计算完了，无需重复计算
			if j-arr[index] >= 0 {
				dp[index][j] = dp[index][j-arr[index]] + dp[index+1][j]
			} else {
				dp[index][j] = dp[index+1][j]
			}
		}
	}
	return dp[0][aim]
}
