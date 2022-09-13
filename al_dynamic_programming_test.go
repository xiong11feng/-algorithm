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

func TestHorseGame1(t *testing.T) {
	x := 3
	y := 3
	k := 10
	actual := HorseGame1(x, y, k)
	expected := 1178562
	assert.Equal(t, expected, actual)
}

func TestHorseGame2(t *testing.T) {
	x := 3
	y := 3
	k := 10
	actual := HorseGame2(x, y, k)
	expected := 1178562
	assert.Equal(t, expected, actual)
}

func BenchmarkHorseGame1(b *testing.B) {
	x := 3
	y := 3
	k := 14
	res := HorseGame1(x, y, k)
	b.Logf("result is %d", res)
}

func BenchmarkHorseGame2(b *testing.B) {
	x := 3
	y := 3
	k := 14
	res := HorseGame2(x, y, k)
	b.Logf("result is %d", res)
}

func TestFindMoney1(t *testing.T) {
	arr := []int{1, 2, 3, 5, 10, 20, 50, 100}
	aim := 5
	actual := FindMoney1(aim, arr)
	expected := 6 //1+1+1+1+1,2+3,1+1+1+2,1+1+3,1+2+2,5
	assert.Equal(t, expected, actual)
}
func TestFindMoney2(t *testing.T) {
	arr := []int{1, 2, 3, 5, 10, 20, 50, 100}
	aim := 5
	actual := FindMoney2(aim, arr)
	expected := 6 //1+1+1+1+1,2+3,1+1+1+2,1+1+3,1+2+2,5
	assert.Equal(t, expected, actual)
}

func TestFndMoney3(t *testing.T) {
	arr := []int{1, 2, 3, 5, 10, 20, 50, 100}
	aim := 5
	actual := FindMoney3(aim, arr)
	expected := 6 //1+1+1+1+1,2+3,1+1+1+2,1+1+3,1+2+2,5
	assert.Equal(t, expected, actual)
}

func BenchmarkFindMoney1(b *testing.B) {
	arr := []int{1, 2, 3, 5, 7, 10, 15, 20, 22, 25, 30, 40, 50, 60, 62, 67, 70, 72, 75, 80, 90, 100}
	aim := 200
	res := FindMoney1(aim, arr)
	b.Logf("result is %d", res)
}

func BenchmarkFindMoney2(b *testing.B) {
	arr := []int{1, 2, 3, 5, 7, 10, 15, 20, 22, 25, 30, 40, 50, 60, 62, 67, 70, 72, 75, 80, 90, 100}
	aim := 50000
	res := FindMoney2(aim, arr)
	b.Logf("result is %d", res)
}

func BenchmarkFindMoney3(b *testing.B) {
	arr := []int{1, 2, 3, 5, 7, 10, 15, 20, 22, 25, 30, 40, 50, 60, 62, 67, 70, 72, 75, 80, 90, 100}
	aim := 50000
	res := FindMoney3(aim, arr)
	b.Logf("result is %d", res)
}
