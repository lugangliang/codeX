package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 比较重要的一个排序，比冒泡排序效率高，比选择排序稳定
func InsertSort(array []int, len int) {
	for i:=0; i< len; i++ {
		temp := array[i]
		var j int
		for j = i-1; j >= 0; j--{
			// j就是i前面的数，j+1就是i
			if array[j] > temp {
				array[j+1] = array[j]
			} else {
				break
			}
		}
		// j == -1，说明元素都比array[i]大
		if j == -1 {
			array[0] = temp
		} else {
			// array[j]是第一个比array[i]小的值，所以j+1是array[i]正确的位置
			array[j+1] = temp
		}
	}

	for _, value := range array{
		fmt.Printf("%d\t ",value)

	}
	fmt.Println("")
}

func BubbleSort(array []int, len int) {
	for i:=0; i<len; i++{
		flag := false
		for j:=0 ;j < len -i -1 ; j++ {
			if array[j] > array[j+1]{
				array[j], array[j+1] = array[j+1], array[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	for _, value := range array{
		fmt.Printf("%d\t ",value)

	}
	fmt.Println("")
}

func SelectSort(array []int, len int) {
	for i:=0; i<len; i++ {
		flag := false
		minValueIndex := i
		minValue := array[i]
		for j := i+1; j < len; j++ {
			if array[j] < minValue {
				minValueIndex= j
				minValue = array[j]
				flag = true
			}
		}
		if flag {
			array[i], array[minValueIndex] = array[minValueIndex], array[i]
		}
	}

	for _, value := range array{
		fmt.Printf("%d\t ",value)

	}
	fmt.Println("")
}

func main() {
	for i := 5; i < 300; i++ {
		array := make([]int, i)
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < i; j++ {
			array[j] = rand.Intn(100)
		}

		fmt.Printf("Orignal seraial.\n")
		for k:=0; k<i; k++ {
			fmt.Printf("%d\t", array[k])
		}
		fmt.Printf("\n")

		newSlice1 := make([]int, i)
		copy(newSlice1, array)

		newSlice2 := make([]int, i)
		copy(newSlice2, array)

		newSlice3 := make([]int, i)
		copy(newSlice3, array)

		fmt.Println("Bubble sort")
		BubbleSort(newSlice1, i)

		fmt.Println("Select sort")
		SelectSort(newSlice2, i)

		fmt.Println("Insert sort")
		InsertSort(newSlice3, i)
	}
	fmt.Println("Hello world.")

	return
}

func CheckResult(array []int) {

	i, j:=0,0
	sliceLen := len(array)
	for ; i<sliceLen;i++ {
		for j=i+1; j < sliceLen; j++ {
			if array[i] > array[j]{
				fmt.Println("Sort Error...............")
			}
		}
	}

	return

}