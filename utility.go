package main

import (
	"math/rand"
)

func GenerateRandomArray(maxLength, maxValue int) []int32 {
	arr := make([]int32, rand.Intn(int(maxLength)+1))
	for i := 0; i < len(arr); i++ {
		arr[i] = int32(rand.Intn(maxLength+1) - rand.Intn(maxLength+1))
	}
	return arr
}

func CompareArray(arr1, arr2 []int32) bool {
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

func CopyArray(arr []int32) []int32 {
	arr2 := make([]int32, len(arr))
	for i := 0; i < len(arr); i++ {
		arr2[i] = arr[i]
	}
	return arr2
}
