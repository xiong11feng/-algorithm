package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort_Insert(t *testing.T) {
	items := GenerateRandomArray(1000, 10000)
	items2 := CopyArray(items)
	Sort_Insert(items)
	sort.Slice(items2, func(i, j int) bool {
		return items2[i] < items2[j]
	})
	assert.True(t, CompareArray(items, items2), "")
}

func TestSort_Bubble(t *testing.T) {
	items := GenerateRandomArray(1000, 10000)
	items2 := CopyArray(items)
	Sort_Bubble(items)
	sort.Slice(items2, func(i, j int) bool {
		return items2[i] < items2[j]
	})
	assert.True(t, CompareArray(items, items2), "")
}

func TestSort_Selected(t *testing.T) {
	items := GenerateRandomArray(1000, 10000)
	items2 := CopyArray(items)
	Sort_Selected(items)
	sort.Slice(items2, func(i, j int) bool {
		return items2[i] < items2[j]
	})
	assert.True(t, CompareArray(items, items2), "")
}

func TestSort_Merge(t *testing.T) {
	items := GenerateRandomArray(1000, 10000)
	items2 := CopyArray(items)
	Sort_Merge(items)
	sort.Slice(items2, func(i, j int) bool {
		return items2[i] < items2[j]
	})
	assert.True(t, CompareArray(items, items2), "")
}
