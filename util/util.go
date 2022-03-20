package util

import "fmt"

func checkResult(array []int) {

	i, j := 0, 0
	sliceLen := len(array)
	for ; i < sliceLen; i++ {
		for j = i + 1; j < sliceLen; j++ {
			if array[i] > array[j] {
				fmt.Println("Sort Error...............", array)
			}
		}
	}

	return

}
