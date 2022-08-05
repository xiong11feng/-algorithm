package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapInsert(t *testing.T) {
	//            4
	//          /   \
	//         1     2
	//        / \   / \
	//       6   5 7   3
	items := []int{4, 1, 2, 6, 5, 7, 3}
	HeapInsert(items, 6)
	items2 := []int{4, 1, 3, 6, 5, 7, 2}
	assert.True(t, CompareArray(items, items2), "")
}
