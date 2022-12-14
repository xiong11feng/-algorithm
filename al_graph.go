package main

import (
	"sort"
)

//图，就是点和边的组合
//图的存储方式
//邻接表 - 某节点的直接邻居 A - B，C｜ B - C，D ...
//邻接矩阵  --   ABCD 和 ABCD 组成矩阵

type Graph struct {
	Nodes map[int]*GraphNode
	Edges []*GraphEdge
}

type GraphNode struct {
	Value int          //编号
	In    int          //入度, 有多少边进入这个点
	Out   int          //出度，从这个点有多少边出去
	Nexts []*GraphNode //与自己相临的点
	Edges []*GraphEdge //自己拥有的边
}

func (g *Graph) AddNode(values []int) {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*GraphNode, 0)
	}
	for i := 0; i < len(values); i++ {
		g.Nodes[values[i]] = &GraphNode{Value: values[i]}
	}
}

func (g *Graph) AddEdge(from, to, weight int) bool {

	if _, ok := g.Nodes[from]; !ok {
		return false
	}
	if _, ok := g.Nodes[to]; !ok {
		return false
	}
	edge := &GraphEdge{Weight: weight, From: g.Nodes[from], To: g.Nodes[to]}
	if g.Edges == nil {
		g.Edges = make([]*GraphEdge, 0)
	}
	g.Nodes[from].Out++
	if g.Nodes[from].Nexts == nil {
		g.Nodes[from].Nexts = make([]*GraphNode, 0)
	}
	g.Nodes[from].Nexts = append(g.Nodes[from].Nexts, g.Nodes[to])
	if g.Nodes[from].Edges == nil {
		g.Nodes[from].Edges = make([]*GraphEdge, 0)
	}
	g.Nodes[from].Edges = append(g.Nodes[from].Edges, edge)
	g.Nodes[to].In++
	g.Edges = append(g.Edges, edge)
	return true
}

type GraphEdge struct {
	Weight int
	From   *GraphNode
	To     *GraphNode
}

type MySets struct {
	setMap map[*GraphNode]*[]*GraphNode
}

func (s *MySets) InitMysets(graph *Graph) {
	s.setMap = make(map[*GraphNode]*[]*GraphNode)
	for _, v := range graph.Nodes {
		nodeList := make([]*GraphNode, 0)
		nodeList = append(nodeList, v)
		s.setMap[v] = &nodeList
	}
}

func (s *MySets) isSameSet(node1, node2 *GraphNode) bool {
	return s.setMap[node1] == s.setMap[node2]
}

func (s *MySets) union(from, to *GraphNode) {
	fromSet := s.setMap[from]
	toSet := s.setMap[to]
	for i := 0; i < len(*toSet); i++ {
		*fromSet = append(*fromSet, (*toSet)[i])
		s.setMap[(*toSet)[i]] = fromSet
	}
}

//【图的宽度优先】
//和二叉树相比，二叉树没有环，但是图有环
//使用队列来实现
func BFS_Graph(node *GraphNode, arr *[]int) {
	if node == nil {
		return
	}
	queue := make(Queue, 0)
	set := make(map[*GraphNode]struct{}, 0)
	queue.push(node)
	set[node] = struct{}{}
	for !queue.isEmpty() {
		temp, _ := queue.pop()
		tempNode := temp.(*GraphNode)
		*arr = append(*arr, tempNode.Value)
		if tempNode.Nexts != nil {
			for _, node1 := range tempNode.Nexts {
				if _, ok := set[node1]; !ok {
					queue.push(node1)
					set[node1] = struct{}{}
				}
			}
		}
	}
}

//【图的深度优先遍历】
//使用栈，和hash map
//从源节点开始，一次按深度入栈，然后弹出
//每弹出一个节点，把该节点的下一个没有进过栈的邻节点入栈
//直到栈变空
func DFS_Graph(node *GraphNode, arr *[]int) {
	if node == nil {
		return
	}

	stack := make(Stack, 0)
	set := make(map[*GraphNode]struct{}, 0)
	stack.push(node)
	set[node] = struct{}{}
	*arr = append(*arr, node.Value)
	for !stack.isEmpty() {
		temp, _ := stack.pop()
		tempNode := temp.(*GraphNode)
		if _, ok := set[tempNode]; !ok {
			*arr = append(*arr, tempNode.Value)
			set[tempNode] = struct{}{}
		}
		if tempNode.Nexts != nil {
			for _, node1 := range tempNode.Nexts {
				if _, ok := set[node1]; !ok {
					stack.push(tempNode)
					stack.push(node1)
					break
				}
			}
		}
	}

	//另一种思路
	// stack := make(Stack, 0)
	// set := make(map[*GraphNode]struct{}, 0)
	// stack.push(node)
	// for !stack.isEmpty() {
	// 	temp, _ := stack.pop()
	// 	tempNode := temp.(*GraphNode)
	// 	if _, ok := set[tempNode]; !ok {
	// 		*arr = append(*arr, tempNode.Value)
	// 		set[tempNode] = struct{}{}
	// 		if tempNode.Nexts != nil {
	// 			for _, node1 := range tempNode.Nexts {
	// 				if _, ok := set[node1]; !ok {
	// 					stack.push(node1)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}

//【拓扑排序】 https://zhuanlan.zhihu.com/p/339709006
//在图论中，拓扑排序（Topological Sorting）是一个有向无环图（DAG, Directed Acyclic Graph）的所有顶点的线性序列。且该序列必须满足下面两个条件：
//1.每个顶点出现且只出现一次。
//2.若存在一条从顶点 A 到顶点 B 的路径，那么在序列中顶点 A 出现在顶点 B 的前面。
func TopSort(graph *Graph, arr *[]int) {
	inMap := make(map[*GraphNode]int)
	zeorInQueue := make(Queue, 0)

	for _, value := range graph.Nodes {
		if value.In == 0 {
			zeorInQueue.push(value)
		} else {
			inMap[value] = value.In
		}
	}
	for !zeorInQueue.isEmpty() {
		temp, _ := zeorInQueue.pop()
		tempNode := temp.(*GraphNode)
		*arr = append(*arr, tempNode.Value)

		if tempNode.Nexts != nil {
			for i := 0; i < len(tempNode.Nexts); i++ {
				inMap[tempNode.Nexts[i]]--
				if inMap[tempNode.Nexts[i]] == 0 {
					zeorInQueue.push(tempNode.Nexts[i])
				}
			}
		}
	}
}

//【Kruska 算法】无向图，生成最小生成树
//最小生成树, 保证连通性，并且权值最小
//算法介绍：去掉所有边，从最短边开始，逐渐增加，增加后没有形成环，增加，形成环了舍弃
func KruskaSmallTree(graph *Graph) []GraphEdge {
	sort.Slice(graph.Edges, func(i, j int) bool {
		return graph.Edges[i].Weight < graph.Edges[j].Weight
	})
	sortEdegeQueue := make(Queue, 0)
	for _, v := range graph.Edges {
		sortEdegeQueue.push(v)
	}
	mySet := MySets{}
	mySet.InitMysets(graph)
	result := make([]GraphEdge, 0)
	for !sortEdegeQueue.isEmpty() {
		cur, _ := sortEdegeQueue.pop()
		curEdge := cur.(*GraphEdge)
		if !mySet.isSameSet(curEdge.From, curEdge.To) {
			mySet.union(curEdge.From, curEdge.To)
			result = append(result, *curEdge)
		}
	}
	return result
}

//【Prime 算法】无向图，生成最小生成树
func PrimeSmallTree(graph *Graph) []GraphEdge {
	sortEdegeQueue := make([]*GraphEdge, 0)
	// for _, v := range graph.Edges {
	// 	sortEdegeQueue = append(sortEdegeQueue, v)
	// }
	alreadHandlSet := make(map[*GraphNode]struct{})
	result := make([]GraphEdge, 0)
	for _, v := range graph.Nodes {
		if _, ok := alreadHandlSet[v]; !ok {
			alreadHandlSet[v] = struct{}{}
			sortEdegeQueue = append(sortEdegeQueue, v.Edges...)
			for len(sortEdegeQueue) > 0 {
				minEdge := PrimeSmallTree_findSmallEdge(&sortEdegeQueue)
				if _, ok := alreadHandlSet[minEdge.To]; !ok {
					alreadHandlSet[minEdge.To] = struct{}{}
					result = append(result, *minEdge)
					sortEdegeQueue = append(sortEdegeQueue, minEdge.To.Edges...)
				}
			}
		}
	}
	return result
}

func PrimeSmallTree_findSmallEdge(edges *[]*GraphEdge) *GraphEdge {
	min := (*edges)[0].Weight
	index := 0
	for idx, v := range *edges {
		if min >= v.Weight {
			min = v.Weight
			index = idx
		}
	}
	minEdge := (*edges)[index]
	temp := make([]*GraphEdge, 0)
	temp = append(temp, (*edges)[:index]...)
	temp = append(temp, (*edges)[index+1:]...)
	*edges = temp
	return minEdge
}

//【Dijkstra 算法】 使用范围，没有权值为负的边
// https://zhuanlan.zhihu.com/p/338414118
// Dijkstra 算法是一个基于「贪心」、「广度优先搜索」、「动态规划」求一个图中一个点到其他所有点的最短路径的算法，时间复杂度 O(n2)
func Dijkstra_Graph(head *GraphNode) map[*GraphNode]int {
	//distanceMap 从 head出发，到key点的最小距离
	//没有key，表示正无穷
	distanceMap := make(map[*GraphNode]int)

	//先将自己加入节点
	distanceMap[head] = 0

	//已经处理过的节点，无需再处理
	selectedSet := make(map[*GraphNode]struct{})
	minNode := dijkstra_Graph_getMinDsitanceAndSelectedNode(distanceMap, selectedSet)
	for minNode != nil {
		//到当前节点需要的值
		currentValue := distanceMap[minNode]
		for _, v := range minNode.Edges {
			toNode := v.To
			if _, ok := distanceMap[toNode]; !ok {
				distanceMap[toNode] = currentValue + v.Weight
			}
			curToNodeValue := distanceMap[toNode]
			if curToNodeValue > currentValue+v.Weight {
				distanceMap[toNode] = currentValue + v.Weight
			}
		}
		//当前节点处理过了，就不再处理
		selectedSet[minNode] = struct{}{}
		minNode = dijkstra_Graph_getMinDsitanceAndSelectedNode(distanceMap, selectedSet)
	}
	return distanceMap
}

func dijkstra_Graph_getMinDsitanceAndSelectedNode(distanceMap map[*GraphNode]int, selectedSet map[*GraphNode]struct{}) *GraphNode {
	min := -1
	var temp *GraphNode
	for k, v := range distanceMap {
		if min == -1 || min > v {
			if _, ok := selectedSet[k]; !ok {
				min = v
				temp = k
			}
		}
	}
	return temp
}
