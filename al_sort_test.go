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

func TestSmallSum(t *testing.T) {
	items := []int{1, 2, 6, 5, 4, 1, 1, 1, 8, 7} //6+10+12+10+8 + 2 + 2 + 2 = 52
	actual := SmallSum(items)
	expected := 52
	assert.Equal(t, expected, actual, "")
}

func TestDutchFlag(t *testing.T) {
	items := []int{1, 2, 6, 5, 4, 1, 1, 1, 8, 7}
	DutchFlag(items, 2)
	expected := int(2)
	assert.Equal(t, expected, items[4], "")
}

func TestSort_Fast_V2(t *testing.T) {
	items := GenerateRandomArray(1000, 10000)
	//items := []int{4, 1, 2, 6, 5, 7, 3}
	items2 := CopyArray(items)
	Sort_Fast_V2(items)
	sort.Slice(items2, func(i, j int) bool {
		return items2[i] < items2[j]
	})
	assert.True(t, CompareArray(items, items2), "")
}
func TestSort_Fast_V3(t *testing.T) {
	items := GenerateRandomArray(1000, 10000)
	//items := []int{4, 1, 2, 6, 5, 7, 3}
	items2 := CopyArray(items)
	Sort_Fast_V3(items)
	sort.Slice(items2, func(i, j int) bool {
		return items2[i] < items2[j]
	})
	assert.True(t, CompareArray(items, items2), "")
}
func TestSort_Heap(t *testing.T) {
	items := GenerateRandomArray(1000, 10000)
	//items := []int{4, 1, 2, 6, 5, 7, 3}
	items2 := CopyArray(items)
	Sort_Heap(items)
	sort.Slice(items2, func(i, j int) bool {
		return items2[i] < items2[j]
	})
	assert.True(t, CompareArray(items, items2), "")
}
