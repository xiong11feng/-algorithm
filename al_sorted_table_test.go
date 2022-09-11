package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsLand(t *testing.T) {
	m := [][]int{{0, 0, 1, 0, 1, 0}, {1, 1, 1, 0, 1, 0}, {1, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0}}
	actual := IsLand(m)
	assert.Equal(t, 3, actual)
}

func generateStr1AndStr2() (str1, str2 string) {
	rand.Seed(time.Now().UnixNano())
	max := 10000000000
	min := 1000000000
	n := rand.Intn(max-min) + min
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		min = 97
		max = 105
		v := rand.Intn(max-min) + min
		sb.WriteByte(byte(v))
	}
	str1 = sb.String()
	sb.Reset()
	max = 5000000
	min = 100000
	n = rand.Intn(max-min) + min
	for i := 0; i < n; i++ {
		min = 97
		max = 105
		v := rand.Intn(max-min) + min
		sb.WriteByte(byte(v))
	}
	str2 = sb.String()
	return
}

func generateStr1(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(max-min) + min
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		min = 97
		max = 105
		v := rand.Intn(max-min) + min
		sb.WriteByte(byte(v))
	}
	return sb.String()
}

func TestKMP(t *testing.T) {
	str1 := "abcdkljalijsdflkjaslkdfjn jknaefjoaiejfoasdflkasdkl"
	str2 := "flkj"
	expected := customIndexOf(str1, str2)
	actual := KMP(str1, str2)
	assert.Equal(t, expected, actual)
}

func BenchmarkKMP(b *testing.B) {
	b.StopTimer()
	str1, str2 := generateStr1AndStr2()
	b.StartTimer()

	time2 := time.Now()
	strings.Index(str1, str2)
	tc2 := time.Since(time2) //计算耗时
	b.Logf("strings index of time cost = %v\n", tc2)

	time3 := time.Now()
	customIndexOf(str1, str2)
	tc3 := time.Since(time3) //计算耗时
	b.Logf("custom strings index of time cost = %v\n", tc3)

	time1 := time.Now()
	KMP(str1, str2)
	tc := time.Since(time1) //计算耗时
	b.Logf("kmp time cost = %v\n", tc)
}

func customIndexOf(s, substr string) int {
	for i := 0; i < len(s); i++ {
		if substr[0] == s[i] {
			temp1 := 0 + 1
			temp2 := i + 1
			for temp1 < len(substr) && temp2 < len(s) && substr[temp1] == s[temp2] {
				temp1++
				temp2++
			}
			if temp1 == len(substr) {
				return i
			}
		}
	}
	return -1
}

func TestManacher(t *testing.T) {
	input := "ababa" //5
	acutal := Manacher(input)
	expected := 5
	assert.Equal(t, expected, acutal)
	input = "abaaba" //6
	acutal = Manacher(input)
	expected = 6
	assert.Equal(t, expected, acutal)
	input = ""
	acutal = Manacher(input)
	expected = 0
	assert.Equal(t, expected, acutal)
	input = "a"
	acutal = Manacher(input)
	expected = 1
	assert.Equal(t, expected, acutal)
	input = "abbbcabcdefedcbaccccab"
	acutal = Manacher(input)
	expected = 13
	assert.Equal(t, expected, acutal)
}

func BenchmarkManacher(b *testing.B) {
	b.StopTimer()
	input := generateStr1(20000, 30000)
	b.StartTimer()
	b.Logf("current result is %d", Manacher(input))
}
