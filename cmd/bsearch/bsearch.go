package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func GenerateData(n int) []int {
	result := make([]int, n)
	time.Sleep(100 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < n; j++ {
		result[j] = rand.Intn(100)
	}

	result = MergeSort(result, false)

	return result
}

// 在数组中查找第一个大于等于value的值
func Bsearch1(slice []int, value int) int {
	low,mid,high := 0,0,0
	high = len(slice) - 1

	for ; low <= high; {
		mid = low + (high-low)>>1
		if slice[mid] < value {
			low = mid + 1
		} else if slice[mid] >= value{
			if mid == 0 || slice[mid-1] < value {
				return mid
			}
			high = mid - 1
		}
	}

	return -1
}

// 在数组中查找最后一个小于等于value的值
func Bsearch2(slice []int, value int) int {
	low, mid, high := 0,0,0
	high = len(slice) - 1

	for ;low <= high; {
		mid = low + (high - low )>>1
		if slice[mid] > value {
			high = mid - 1
		} else if slice[mid] <= value {
			if mid == slice[ len(slice) -1 ] || slice[mid+1] > value {
				return mid
			}
			low = mid + 1
		}
 	}

	return 0
}

func main(){

	result := GenerateData(10)

	fmt.Println("origin : ", result)

	fmt.Println(Bsearch1(result, 33))

	fmt.Println(Bsearch2(result, 44))

	fmt.Println("hello world.")
	return
}
