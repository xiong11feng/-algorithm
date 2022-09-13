//有序表
//所有操作都是logN，高效
//实现有效表的方式 ，BST平衡搜索二叉树（红黑树，AVL，SB），跳表
package main

//【岛问题】 一个矩阵，0，1组成，左右上下相连的1，组成岛，求岛的数量
//001010
//111010
//100100
//000000
//上边有三个岛
//思路：从头开始，找到1，infection（感染），将连着1改成2
//并行思路：使用并查集

func IsLand(m [][]int) int {
	if len(m) == 0 {
		return 0
	}
	res := 0
	x := len(m)
	y := len(m[0])
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if m[i][j] == 1 {
				res++
				isLandInfection(m, i, j, x, y)
			}
		}
	}
	return res
}

func isLandInfection(m [][]int, i, j, x, y int) {
	if i >= 0 && i < x && j >= 0 && j < y {
		if m[i][j] == 1 {
			//将1改成2，递归到上下左右
			m[i][j] = 2

			isLandInfection(m, i-1, j, x, y) //上
			isLandInfection(m, i+1, j, x, y) //下
			isLandInfection(m, i, j-1, x, y) //左
			isLandInfection(m, i, j+1, x, y) //右
		}
	}
}

//并查集
//查询两个集合是不是有同一个元素
//合并两个集合
//都是O(1)
//代码实现：https://github.com/spakin/disjoint

//[KMP]算法
//判断 str2是不是str1的子串
//暴力解法：一个一个比较O(n*m)
//给 str2 增加一个最长前缀数组
//详解：https://zhuanlan.zhihu.com/p/145536254
//例如str2 = "ababcababc",数组是[-1,0,0,2,0,0,3,]

func KMP(str1, str2 string) int {
	sa1 := []byte(str1)
	sa2 := []byte(str2)
	i1 := 0 //str1中比对的位置
	i2 := 0 //str2中比对的位置

	next := kmpGetNextArray(sa2) //O(M)

	for i1 < len(str1) && i2 < len(str2) {
		if sa1[i1] == sa2[i2] {
			i1++
			i2++
		} else if next[i2] == -1 { //i2==0,i2在0位置，
			i1++
		} else { //i2还能往前跳
			i2 = next[i2]
		}
	}
	if i2 == len(str2) {
		return i1 - len(str2)
	}
	return -1
}

func kmpGetNextArray(a []byte) []int {
	next := make([]int, len(a))
	if len(a) == 1 {
		next[0] = -1 //人为规定-1
	}
	next[0] = -1 //人为规定-1
	next[1] = 0  //人为规定0
	i := 2       //next 数组的位置
	cn := 0      //和i-1字符比较，代表当前使用的信息是多少
	for i < len(a) {
		if a[cn] == a[i-1] {
			next[i] = cn + 1
			i++
			cn++
		} else if cn > 0 { //cn ==0到头了，不需要向前调整cn了
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}

//[manacher 算法]
//manacher : https://www.geeksforgeeks.org/manachers-algorithm-linear-time-longest-palindromic-substring-part-1/
//求字符串中的最大回文长度
//“abccb” = 4
//“dabac” = 3
func Manacher(str string) int {
	if str == "" {
		return 0
	}
	newMancherBytes := manacherString(str)
	rArr := make([]int, len(newMancherBytes)) //回文半径数组
	c := -1                                   //最右回文边界的中心
	r := -1                                   //最右回文边界的下一个位置
	max := 0
	for i := 0; i < len(newMancherBytes); i++ {
		if i >= r { //如果当前的i在r外，当前位置不用验证，向两边暴力验证
			rArr[i] = 1
		} else if rArr[2*c-i] > r-i {
			rArr[i] = r - i
		} else {
			rArr[i] = rArr[2*c-i]
		}

		//上边求出的是i至少不用验证的大小
		//1.如果i再r之外，只能进行暴力左右比较，没法加速，
		//2.如果i在r之内，有三种情况
		//2.1 [...[  ( i' )    c   ( i )    ]...]  i‘ 的半径，在c的半径内，i的回味半径值等于i’
		//2.2 [..（.[ i' )    c        ( i ].）..]  i‘ 的半径 在c的半径外，i的回文半径等 r-i
		//2.3　如果i‘ 的半径 ，c的外半径恰巧重合，那i的回文半径至少是r-1的位置，后续只能暴力比较
		//后续同时左右验证

		for i+rArr[i] < len(newMancherBytes) && i-rArr[i] >= 0 {
			if newMancherBytes[i+rArr[i]] == newMancherBytes[i-rArr[i]] {
				rArr[i]++
				continue
			}
			break
		}
		if i+rArr[i] > r {
			r = i + rArr[i]
			c = i
		}
		if rArr[i] > max {
			max = rArr[i]
		}
	}
	//处理串最大半径长度减1，是原串的直径
	return max - 1
}

func manacherString(str string) []byte {
	newL := 2*len(str) + 1
	result := make([]byte, newL)
	result[0] = '#'
	for i := 0; i < len(str); i++ {
		result[2*i+1] = str[i]
		result[2*i+2] = '#'
	}
	return result
}
