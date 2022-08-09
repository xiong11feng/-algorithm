package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMidNode(t *testing.T) {
	Node1 := &SingalNode{Value: 1}
	Node2 := &SingalNode{Value: 2}
	Node1.Next = Node2
	Node3 := &SingalNode{Value: 3}
	Node2.Next = Node3
	Node4 := &SingalNode{Value: 4}
	Node3.Next = Node4
	Node5 := &SingalNode{Value: 5}
	Node4.Next = Node5
	actual := FindMidNode(Node1)
	assert.Equal(t, 3, actual.Value)
	Node6 := &SingalNode{Value: 6}
	Node5.Next = Node6
	actual = FindMidNode(Node1)
	assert.Equal(t, 4, actual.Value)
}

func TestReversalNode(t *testing.T) {
	Node1 := &SingalNode{Value: 1}
	Node2 := &SingalNode{Value: 2}
	Node1.Next = Node2
	Node3 := &SingalNode{Value: 3}
	Node2.Next = Node3
	Node4 := &SingalNode{Value: 4}
	Node3.Next = Node4
	Node5 := &SingalNode{Value: 5}
	Node4.Next = Node5
	Node6 := &SingalNode{Value: 6}
	Node5.Next = Node6
	actual := ReversalNode(Node1)
	assert.Equal(t, 6, actual.Value)
	assert.Equal(t, 5, actual.Next.Value)
	assert.Equal(t, 4, actual.Next.Next.Value)
	assert.Equal(t, 3, actual.Next.Next.Next.Value)
	assert.Equal(t, 2, actual.Next.Next.Next.Next.Value)
	assert.Equal(t, 1, actual.Next.Next.Next.Next.Next.Value)
}

func TestIsSymmetryLinkedList(t *testing.T) {
	Node1 := &SingalNode{Value: 1}
	Node2 := &SingalNode{Value: 2}
	Node1.Next = Node2
	Node3 := &SingalNode{Value: 3}
	Node2.Next = Node3
	Node4 := &SingalNode{Value: 3}
	Node3.Next = Node4
	Node5 := &SingalNode{Value: 2}
	Node4.Next = Node5
	Node6 := &SingalNode{Value: 1}
	Node5.Next = Node6
	actual := IsSymmetryLinkedList(Node1)
	assert.Equal(t, true, actual)
	assert.Equal(t, 1, Node1.Value)
	assert.Equal(t, 2, Node1.Next.Value)
	assert.Equal(t, 3, Node1.Next.Next.Value)
	assert.Equal(t, 3, Node1.Next.Next.Next.Value)
	assert.Equal(t, 2, Node1.Next.Next.Next.Next.Value)
	assert.Equal(t, 1, Node1.Next.Next.Next.Next.Next.Value)

	Node7 := &SingalNode{Value: 1}
	Node6.Next = Node7
	actual = IsSymmetryLinkedList(Node1)
	assert.Equal(t, false, actual)

	node1 := &SingalNode{Value: 1}
	node2 := &SingalNode{Value: 2}
	node1.Next = node2
	node3 := &SingalNode{Value: 3}
	node2.Next = node3
	node5 := &SingalNode{Value: 2}
	node3.Next = node5
	node6 := &SingalNode{Value: 1}
	node5.Next = node6
	actual = IsSymmetryLinkedList(node1)
	assert.Equal(t, true, actual)
	assert.Equal(t, 1, node1.Value)
	assert.Equal(t, 2, node1.Next.Value)
	assert.Equal(t, 3, node1.Next.Next.Value)
	assert.Equal(t, 2, node1.Next.Next.Next.Value)
	assert.Equal(t, 1, node1.Next.Next.Next.Next.Value)
}

func TestListPartition(t *testing.T) {
	Node1 := &SingalNode{Value: 6}
	Node2 := &SingalNode{Value: 2}
	Node1.Next = Node2
	Node3 := &SingalNode{Value: 4}
	Node2.Next = Node3
	Node4 := &SingalNode{Value: 3}
	Node3.Next = Node4
	Node5 := &SingalNode{Value: 5}
	Node4.Next = Node5
	Node6 := &SingalNode{Value: 6}
	Node5.Next = Node6
	actual := ListPartition(Node1, 4)
	assert.Equal(t, 2, actual.Value)
	assert.Equal(t, 3, actual.Next.Value)
	assert.Equal(t, 4, actual.Next.Next.Value)
	assert.Equal(t, 6, actual.Next.Next.Next.Value)
	assert.Equal(t, 5, actual.Next.Next.Next.Next.Value)
	assert.Equal(t, 6, actual.Next.Next.Next.Next.Next.Value)
}
