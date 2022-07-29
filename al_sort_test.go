package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort_Insert(t *testing.T) {
	items := []uint32{6, 3, 2, 1, 10, 8, 9, 2, 4, 4, 4, 30, 20, 7}
	Sort_Insert(items)
	actual := arrayTostring(items)
	assert.Equal(t, "12234446789102030", actual)
}

func TestSort_Bubble(t *testing.T) {
	items := []uint32{6, 3, 2, 1, 10, 8, 9, 2, 4, 4, 4, 30, 20, 7}
	Sort_Bubble(items)
	actual := arrayTostring(items)
	assert.Equal(t, "12234446789102030", actual)
}

func TestSort_Selected(t *testing.T) {
	items := []uint32{6, 3, 2, 1, 10, 8, 9, 2, 4, 4, 4, 30, 20, 7}
	Sort_Selected(items)
	actual := arrayTostring(items)
	assert.Equal(t, "12234446789102030", actual)
}
