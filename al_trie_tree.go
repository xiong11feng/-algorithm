package main

//前缀树，即字典树，又称单词查找树或键树，是一种树形结构，是一种哈希树的变种。
//典型应用是用于统计和排序大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计。
//它的优点是：利用字符串的公共前缀来减少查询时间，最大限度地减少无谓的字符串比较。

//前缀树的3个基本性质：
//1.根节点不包含字符，除根节点外每一个节点都只包含一个字符。
//2.从根节点到某一节点，路径上经过的字符连接起来，为该节点对应的字符串。
//3.每个节点的所有子节点包含的字符都不相同。

type TrieNode struct {
	pass  int         //多少个路径经过这个节点
	end   int         //多少个路径以此节点为结束
	nexts []*TrieNode //下一个节点，正常是26个字母，最多有26个子节点，如果字符种类很多，需要用hash表 map[rune]*TrieNode
}

func InitTrieNode() *TrieNode {
	res := &TrieNode{}
	res.nexts = make([]*TrieNode, 26) //26个字母，a,b,c...
	return res
}

func (node *TrieNode) Insert(word string) {
	node.pass++
	temp := node
	for _, v := range word {
		if temp.nexts[v-97] == nil {
			temp.nexts[v-97] = InitTrieNode()
		}
		temp = temp.nexts[v-97]
		temp.pass++
	}
	temp.end++
}

//删除字符串
//1.检查是否word存在，用search
//2.再删除，沿途pass--，最后end--
func (node *TrieNode) Delete(word string) {
	if node.Search(word) > 0 {
		temp := node
		temp.pass--
		for _, v := range word {
			if temp.nexts[v-97].pass == 1 {
				temp.nexts[v-97] = nil
				return
			}
			temp = temp.nexts[v-97]
			temp.pass--
		}
		temp.end--
	}
}

func (node *TrieNode) Search(word string) int {
	temp := node
	for _, v := range word {
		temp = temp.nexts[v-97]
		if temp == nil {
			return 0
		}
	}
	return temp.end
}

//有多少字符串，是以word作为前缀的
func (node *TrieNode) PrefixNumber(word string) int {
	temp := node
	for _, v := range word {
		temp = temp.nexts[v-97]
		if temp == nil {
			return 0
		}
	}
	return temp.pass
}

//【贪心算法】
//贪心算法，又名贪婪法，是寻找最优解问题的常用方法，
//这种方法模式一般将求解过程分成若干个步骤，但每个步骤都应用贪心原则，选取当前状态下最好/最优的选择（局部最有利的选择），并以此希望最后堆叠出的结果也是最好/最优的解

//【会议室安排问题】
//1。哪个会议结束时间早，就先安排哪个会议。
//2。之后将不能安排的去掉
//3.重复1，2

//【霍夫曼编码】
//金条切割，每次切需要花费切割今天长度对应的铜板数
//例如长度60的金条，现需要10，20，30三根金条，
//切割方法1: 可以60 -> {30,30}, 30 -> {10,20} 花费90
//切割方法2: 60 ->{10,50},50->{20,30} 花费110
//算法：使用小根堆，取前两个，加和，然后将和推进小根堆 【霍夫曼编码】

//【N皇后问题】
// n*n 的方格中，摆放在每行一个皇后，每一行，每一列，每一斜线，都不能有两个皇后，共有多少种摆放的方法
// 思路：先放置第一行的一个皇后，第一行确定后，找到所有可以放置第二行的点，确定第二行，开始第三行
func NQueue(n int) int {
	i := 0
	arr := make([]*int, n)
	return nQueuePrcess_1(i, n, &arr)
}

//第i行，返回后面的行有多少种摆放方法
func nQueuePrcess_1(i, n int, arr *[]*int) int {
	//如果 i=n 表示前n行都摆放完成
	if i == n {
		return 1
	}
	sum := 0
	for idx := 0; idx < n; idx++ {
		if nQueuePrcess_1_checkIdx(i, idx, arr) {
			(*arr)[i] = &idx
			sum += nQueuePrcess_1(i+1, n, arr)
			(*arr)[i] = nil
		}
	}
	return sum
}

func nQueuePrcess_1_checkIdx(i, idx int, arr *[]*int) bool {
	for j := 0; j < i; j++ {
		//同列
		if *(*arr)[j] == idx {
			return false
		}
		//同斜线,绝对值
		t1 := i - j
		if t1 < 0 {
			t1 = -t1
		}
		t2 := idx - *(*arr)[j]
		if t2 < 0 {
			t2 = -t2
		}
		if t1 == t2 {
			return false
		}
	}
	return true
}

//使用位运算，加快运算时间
//思路，每一行有n个二进制表示，通过位运算计算可用的位置
//计算限制
func NQueue_2(n int) int {
	//限制，左边全零，右边n个1
	limit := 1<<n - 1
	return nQueue_2(limit, 0, 0, 0) //第一次没有任何限制

}

//总限制，列限制，左斜线限制，右斜线限制，1表示不能填，0表示可以填
func nQueue_2(limit, colLimit, leftDiaLimit, rightDimLimit int) int {
	//所有的列，全都被限制了，说明已经放满了，因为每次只填一个1
	if limit == colLimit {
		return 1
	}

	result := 0
	mostRightOne := 0
	//三个限制或，取到所有的不能填的1，取返，0表示不能填，1表示能填，再和limit 求与
	//求出，所有能填的位置，1表示可以填的位置
	pos := limit & (^(colLimit | leftDiaLimit | rightDimLimit))
	for pos > 0 {
		//最右侧的1的取法，pos取返+1，再与上Pos
		mostRightOne = pos & (^pos + 1)
		pos ^= mostRightOne
		result += nQueue_2(limit, colLimit|mostRightOne,
			leftDiaLimit<<1|mostRightOne<<1,
			rightDimLimit>>1|mostRightOne>>1)
	}
	return result

}
