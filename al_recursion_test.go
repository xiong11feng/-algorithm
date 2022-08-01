package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxValue(t *testing.T) {
	items := GenerateRandomArray(20000, 10000)
	expected := items[0]
	for i := 0; i < len(items); i++ {
		if expected < items[i] {
			expected = items[i]
		}
	}
	actual := GetMaxValue(items)

	assert.Equal(t, expected, actual, "")
}
