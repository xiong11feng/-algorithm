package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//无向图
//         1
//      /  |  \
//     2 - 3 - 4
//         | /
//          5
func GetTestGraphNode() *GraphNode {
	node1 := &GraphNode{Value: 1}
	node1.Nexts = make([]*GraphNode, 0)
	node2 := &GraphNode{Value: 2}
	node2.Nexts = make([]*GraphNode, 0)
	//2->1
	node2.Nexts = append(node2.Nexts, node1)
	node3 := &GraphNode{Value: 3}
	node3.Nexts = make([]*GraphNode, 0)
	//3->1
	node3.Nexts = append(node2.Nexts, node1)
	node4 := &GraphNode{Value: 4}
	node4.Nexts = make([]*GraphNode, 0)
	//4->1
	node4.Nexts = append(node2.Nexts, node1)
	//1->2
	node1.Nexts = append(node1.Nexts, node2)
	//1->3
	node1.Nexts = append(node1.Nexts, node3)
	//1->4
	node1.Nexts = append(node1.Nexts, node4)

	//2->3
	node2.Nexts = append(node2.Nexts, node3)
	//3->2
	node3.Nexts = append(node2.Nexts, node2)
	//3->4
	node3.Nexts = append(node2.Nexts, node4)
	//4->3
	node4.Nexts = append(node2.Nexts, node3)

	node5 := &GraphNode{Value: 5}
	node5.Nexts = make([]*GraphNode, 0)
	//5->3
	node5.Nexts = append(node5.Nexts, node3)
	//5->4
	node5.Nexts = append(node5.Nexts, node4)

	//3->5
	node3.Nexts = append(node5.Nexts, node5)
	//4->5
	node4.Nexts = append(node5.Nexts, node5)

	return node1
}

//      3
//   >  > \
//  /   |   >
// 1 -> 2 -> 4 -> 5 -> 6
//
//
func GetTopGraph() *Graph {
	g := &Graph{}
	g.AddNode([]int{1, 2, 3, 4, 5, 6})
	g.AddEdge(1, 2, 100)
	g.AddEdge(1, 3, 100)
	g.AddEdge(2, 3, 100)
	g.AddEdge(2, 4, 100)
	g.AddEdge(3, 4, 100)
	g.AddEdge(4, 5, 100)
	g.AddEdge(5, 6, 100)
	return g
}

//      3
//  /   |   \
// 1 -  2 -  4 -  5 -  6
//
// 无向图
func GetTopGraphWithWeight() *Graph {
	g := &Graph{}
	g.AddNode([]int{1, 2, 3, 4, 5, 6})
	g.AddEdge(1, 2, 10)
	g.AddEdge(2, 1, 10)
	g.AddEdge(1, 3, 20)
	g.AddEdge(3, 1, 20)
	g.AddEdge(2, 3, 30)
	g.AddEdge(3, 2, 30)
	g.AddEdge(2, 4, 40)
	g.AddEdge(4, 2, 40)
	g.AddEdge(3, 4, 50)
	g.AddEdge(4, 3, 50)
	g.AddEdge(4, 5, 60)
	g.AddEdge(5, 4, 60)
	g.AddEdge(5, 6, 70)
	g.AddEdge(6, 5, 70)
	return g
}

func TestBFS_Graph(t *testing.T) {
	gNode := GetTestGraphNode()
	arr := make([]int, 0)
	BFS_Graph(gNode, &arr)
	actual := getArrString(arr)
	assert.Equal(t, "12345", actual)
}

func TestDFS_Graph(t *testing.T) {
	gNode := GetTestGraphNode()
	arr := make([]int, 0)
	DFS_Graph(gNode, &arr)
	actual := getArrString(arr)
	assert.Equal(t, "12345", actual)
}

func TestTopSort(t *testing.T) {
	gNode := GetTopGraph()
	arr := make([]int, 0)
	TopSort(gNode, &arr)
	actual := getArrString(arr)
	assert.Equal(t, "123456", actual)
}

func TestKruskaSmallTree(t *testing.T) {
	gNode := GetTopGraphWithWeight()
	edges := KruskaSmallTree(gNode)
	actual := 0
	for i := 0; i < len(edges); i++ {
		actual += edges[i].Weight
	}
	assert.Equal(t, 200, actual)
}

func TestPrimeSmallTree(t *testing.T) {
	gNode := GetTopGraphWithWeight()
	edges := PrimeSmallTree(gNode)
	actual := 0
	for i := 0; i < len(edges); i++ {
		actual += edges[i].Weight
	}
	assert.Equal(t, 200, actual)
}
