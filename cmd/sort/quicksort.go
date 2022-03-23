package main

import (
	"codeX/util"
	"fmt"
	"math/rand"
	"time"
)

/*
* quick_sort(p,r) = quick_sort(p, pivot) + quick_sort(pivot, r)
* 终止条件 : p >= r
*/

func PartitionLastData(slice []int, start, end int, log bool) int {
	// 将slice分成已排序区间和未排序区间
	// i指向已排序区间的最后一个值，j指向未排序区间的第一个值

	if log {
		fmt.Println("before partion : ", slice)
	}

	i:=start

	pivot := slice[end]

	for j:=start; j < end; j++ {
		if slice[j] < pivot  {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}


	slice[i], slice[end] = slice[end], slice[i]
	if log {
		fmt.Println(slice)
	}

	if log {
		fmt.Println("after partion :  ", slice)
	}
	return i
}

func QuickSort(slice []int, start, end int) {
	// 什么时候等于，[10,11,12]，分区后，pivot返回2，start=3,end=2  QuickSort(slice, pivot+1, end)
	// 什么时候大于，[20,19] 分区后，pivot返回0，start=0，end=0-1   QuickSort(slice, start, pivot-1)
	if start >= end {
		fmt.Println("quick sort start >= end   : ", slice, start, end)
		return
	}

	fmt.Printf("before partion slice %v start %d end %d\n",
		slice[start:end+1], start ,end )

	pivot := PartitionLastData(slice, start, end, false)

	fmt.Printf("after partion, slice1 : %v, pivot_value %d slice2 : %v pivot %d \n",
				slice[start:pivot], slice[pivot], slice[pivot+1:end+1], pivot)

	// 分成了三段，小于pivot，pivot，大于pivot。
	QuickSort(slice, start, pivot-1)
	QuickSort(slice, pivot+1, end)
}

func main() {
	for i := 6; i < 7; i++ {
		array := make([]int, i)
		time.Sleep(1*time.Second)
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < i; j++ {
			array[j] = rand.Intn(100)
		}

		//	27 93 43 44 27
		//array = []int{11, 4, 80, 73}
		fmt.Println("before sort", array)
		QuickSort(array, 0, len(array)-1)
		util.CheckResult(array)
		fmt.Println("after sort", array)
	}
	fmt.Println("Hello world.")

	return
}
