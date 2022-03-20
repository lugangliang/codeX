package main

import (
"fmt"
"math/rand"
"time"
)

/*
*  MergeSort(p, r) = merge(MergeSort(p, q), mergeSort(r+1,r))
*  p >= r 终止
*/

func merge(array1 []int, array2 []int, log bool) []int{
	if log{
		fmt.Println("merge start : ", array1, array2)

	}

	len1, len2 := len(array1), len(array2)

	result := make([]int, len1+len2)

	i,j,k := 0,0,0

	for ;i < len1 && j < len2; k++{
		if array1[i] <= array2[j] {
			result[k] = array1[i]
			i++
		} else {
			result[k] = array2[j]
			j++
		}
	}

	if i == len1 && j <= len2 {
		copy(result[k:], array2[j:])
	}

	if j == len2 && i != len1 {
		copy(result[k:], array1[i:])
	}

	if log{
		fmt.Println("merge result : ", result)

	}
	return result
}


func MergeSort(origin []int, log bool) []int{
	if log{
		fmt.Printf("divide orign %v \n", origin)
	}
	if len(origin) == 1 {
		return  origin
	}
	var middle  = len(origin)/2

	slice1 := origin[:middle]
	slice2 := origin[middle:]

	if log{
		fmt.Println("slice1 , slice2 ,", slice1, slice2)
	}

	return merge(MergeSort(slice1, false),	MergeSort(slice2, false), false)
}

func checkResult(array []int) {

	i, j:=0,0
	sliceLen := len(array)
	for ; i<sliceLen;i++ {
		for j=i+1; j < sliceLen; j++ {
			if array[i] > array[j]{
				fmt.Println("Sort Error...............", array)
			}
		}
	}

	return

}

func main() {
	for i := 4; i < 70; i++ {
		array := make([]int, i)
		time.Sleep(100*time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < i; j++ {
			array[j] = rand.Intn(100)
		}

		fmt.Printf("Orignal sort. %v \n", array)

		//fmt.Println(MergeSort(array))
		result := MergeSort(array, false)
		checkResult(result)
		fmt.Printf("after  sort. %v \n", result)
	}

	fmt.Println("Hello world.")

	return
}