package main

import (
	"math/rand"
	"testing"
	"time"

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

func TestFindAllPermutation(t *testing.T) {
	result := make([]string, 0)
	FindAllPermutation("abcdef", &result)
	actual := len(result)
	expected := 6 * 5 * 4 * 3 * 2
	assert.Equal(t, expected, actual, "")
}

func TestLRArrayGame(t *testing.T) {
	arr := []int{1, 2, 100, 4, 6, 4}
	actual1 := LRArrayGame(arr)
	expected1 := true
	assert.Equal(t, expected1, actual1, "")
}

func TestReverseStack(t *testing.T) {
	stack := make(Stack, 0)
	rand.Seed(time.Now().UnixNano())
	max := 100
	min := 10
	n := rand.Intn(max-min) + min
	for i := 0; i < n; i++ {
		stack.push(i)
	}

	ReverseStack(&stack)
	for i := 0; i < n; i++ {
		item, _ := stack.pop()
		assert.Equal(t, i, item)
	}
}

func TestCalcIntString2LetterString(t *testing.T) {
	input := "111"
	//"AAA","KA","AK"
	expected := 3
	actual := CalcIntString2LetterString(input)
	assert.Equal(t, expected, actual)
	input = "1112"
	//AAAB,AAL,AKB,KAB,KL
	expected = 5
	actual = CalcIntString2LetterString(input)
	assert.Equal(t, expected, actual)
}
