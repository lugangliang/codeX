package util

import (
	"fmt"
	"math/rand"
	"time"
)

func CheckArraySortResult(array []int) (string, bool) {
	i, j := 0, 0
	sliceLen := len(array)
	for ; i < sliceLen; i++ {
		for j = i + 1; j < sliceLen; j++ {
			if array[i] > array[j] {
				fmt.Println("Sort Error...............", array, array[i], array[j])
				return "Sort Error", false
			}
		}
	}

	return "", true
}

func GenerateData(n int) []int {
	array := make([]int, n)
	time.Sleep(200*time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	for j := 1; j < n; j++ {
		array[j] = rand.Intn(100)
	}

	return array
}

