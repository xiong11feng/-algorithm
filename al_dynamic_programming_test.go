package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobotWalk1(t *testing.T) {
	n := 5
	start := 2
	end := 3
	k := 1
	expected := 1
	actual := RobotWalk1(k, n, start, end)
	assert.Equal(t, expected, actual)

	n = 10
	start = 2
	end = 4
	k = 4
	expected = 4 // 2 3 4 5 4; 2 1 2 3 4; 2 3 2 3 4; 2 3 4 3 4
	actual = RobotWalk1(k, n, start, end)
	assert.Equal(t, expected, actual)
}

func TestRobotWalk2(t *testing.T) {
	n := 5
	start := 2
	end := 3
	k := 1
	expected := 1
	actual := RobotWalk2(k, n, start, end)
	assert.Equal(t, expected, actual)

	n = 10
	start = 2
	end = 4
	k = 4
	expected = 4 // 2 3 4 5 4; 2 1 2 3 4; 2 3 2 3 4; 2 3 4 3 4
	actual = RobotWalk2(k, n, start, end)
	assert.Equal(t, expected, actual)
}

func BenchmarkRobotWalk1(b *testing.B) {
	n := 100
	start := 30
	end := 42
	k := 35
	RobotWalk1(k, n, start, end)
}

func BenchmarkRobotWalk2(b *testing.B) {
	n := 100
	start := 30
	end := 42
	k := 35
	RobotWalk2(k, n, start, end)
}
