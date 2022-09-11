package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitAdd(t *testing.T) {
	a := 110
	b := 4000
	actual := BitAdd(a, b)
	expected := a + b
	assert.Equal(t, expected, actual)

	a = -110
	b = -4000
	actual = BitAdd(a, b)
	expected = a + b
	assert.Equal(t, expected, actual)

	a = -100
	b = 4000
	actual = BitAdd(a, b)
	expected = a + b
	assert.Equal(t, expected, actual)
}

func TestBitNegNumber(t *testing.T) {
	a := 10
	actual := BitNegNumber(a)
	expected := -a
	assert.Equal(t, expected, actual)

	a = -10
	actual = BitNegNumber(a)
	expected = -a
	assert.Equal(t, expected, actual)
}

func TestBitMinus(t *testing.T) {
	a := 110
	b := 4000
	actual := BitMinus(a, b)
	expected := a - b
	assert.Equal(t, expected, actual)

	a = 110
	b = -4000
	actual = BitMinus(a, b)
	expected = a - b
	assert.Equal(t, expected, actual)

	a = 110
	b = 100
	actual = BitMinus(a, b)
	expected = a - b
	assert.Equal(t, expected, actual)
}

func TestBitMulti(t *testing.T) {
	a := 10
	b := 10
	expected := a * b
	actual := BitMulti(a, b)
	assert.Equal(t, expected, actual)

	a = -10
	b = 198
	expected = a * b
	actual = BitMulti(a, b)
	assert.Equal(t, expected, actual)

	a = 100
	b = -1980
	expected = a * b
	actual = BitMulti(a, b)
	assert.Equal(t, expected, actual)
}

func TestBitDiv(t *testing.T) {
	a := 10
	b := 10
	expected := a / b
	actual := BitDiv(a, b)
	assert.Equal(t, expected, actual)

	a = 21
	b = 10
	expected = a / b
	actual = BitDiv(a, b)
	assert.Equal(t, expected, actual)

	a = 31
	b = -10
	expected = a / b
	actual = BitDiv(a, b)
	assert.Equal(t, expected, actual)

}
