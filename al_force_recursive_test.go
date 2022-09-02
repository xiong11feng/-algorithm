package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHanoi(t *testing.T) {
	for i := 1; i < 10; i++ {
		actual := Hanoi(i)
		expected := 1<<i - 1
		assert.Equal(t, expected, actual, "")
	}
}

func TestSubSequenceString(t *testing.T) {
	result := make([]string, 0)
	SubSequenceString("abc", &result)
	actual := len(result)
	expected := 8
	assert.Equal(t, expected, actual, "")
}
