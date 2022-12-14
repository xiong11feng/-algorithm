package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXor_ExchangeItem(t *testing.T) {

	items := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	Xor_ExchangeItem(items, 0, 8)
	assert.Equal(t, "8123456709", arrayTostring(items))

	Xor_ExchangeItem(items, 2, 3)
	assert.Equal(t, "8132456709", arrayTostring(items))
}

func arrayTostring(items []int) string {
	value := ""
	for i := 0; i < len(items); i++ {
		value += fmt.Sprintf("%d", items[i])
	}
	return value
}

func TestXor_FindOddDights(t *testing.T) {
	items := []int{1, 2, 2, 3, 3, 4, 4, 4, 4}
	actual := Xor_FindOddDights(items)
	assert.Equal(t, int(1), actual)

	items = []int{2}
	actual = Xor_FindOddDights(items)
	assert.Equal(t, int(2), actual)
}

func TestXor_Xor_FindTwoOddDights(t *testing.T) {
	items := []int{1, 2, 2, 3, 3, 4, 4, 4, 4, 5}
	a, b := Xor_FindTwoOddDights(items)
	assert.Equal(t, int(6), a+b)

	items = []int{2, 22, 22, 33, 44, 44, 44, 44, 55, 55, 66, 66, 66, 66}
	a, b = Xor_FindTwoOddDights(items)
	assert.Equal(t, int(35), a+b)
}
