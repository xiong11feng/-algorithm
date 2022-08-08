package main

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(maxLength, maxValue int) []int {
	rand.Seed(time.Hour.Nanoseconds())
	arr := make([]int, rand.Intn(int(maxLength)+1))
	for i := 0; i < len(arr); i++ {
		arr[i] = int(rand.Intn(maxLength+1) - rand.Intn(maxLength+1))
	}
	return arr
}

func CompareArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func CopyArray(arr []int) []int {
	arr2 := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		arr2[i] = arr[i]
	}
	return arr2
}
