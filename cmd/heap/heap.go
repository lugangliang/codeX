package main

import (
	"codeX/util"
	"fmt"
	"strconv"
)

type Heap struct {
	slice []int
	cap int
	count int
}

func (heap *Heap) String() string {
	var a string

	for i:=1; i<=heap.count; i++ {
		a += " " + strconv.Itoa(heap.slice[i])
	}

	return a
}



func (heap *Heap) Printf() {
	fmt.Println(heap.slice[1:heap.count+1])
}

func (heap *Heap) Insert(value int) {
	if heap.count >= heap.cap {
		return
	}

	heap.count++
	i := heap.count
	heap.slice[i] = value

	// 循环的终止条件是i/2，表示i的父节点至少是根节点
	for ;i/2 >= 1; {
		if heap.slice[i] > heap.slice[i/2] {
			heap.slice[i], heap.slice[i/2] = heap.slice[i/2], heap.slice[i]
			i = i/2
		} else {
			break
		}
	}

	return
}

func (heap *Heap) RemoveMaxElement() int {
	if heap.count == 0 {
		return 0
	}

	max := heap.slice[1]

	heap.slice[1] = heap.slice[heap.count]

	heap.count--

	i := 1

	for ;; {
		maxPos := i
		if 2*i <= heap.count && heap.slice[i] < heap.slice[2*i] {
			maxPos = 2 * i
		}

		if 2*i+1 <= heap.count && heap.slice[maxPos] < heap.slice[2*i+1] {
			maxPos = 2*i+1
		}

		if maxPos != i {
			heap.slice[i], heap.slice[maxPos] = heap.slice[maxPos], heap.slice[i]

			i = maxPos
		} else {
			break
		}

	}

	return max
}

func (heap *Heap) GenerateData() {
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(10)
	heap.Insert(8)
	heap.Insert(6)
	heap.Insert(9)
	heap.Insert(7)

	return
}

func heapify(slice []int, n, i int) {
	var maxPos = i
	var leftChild, rightChild int
	for ;  ; {
		leftChild = 2*i
		rightChild = 2*i + 1


		if leftChild <=n && slice[i] < slice[leftChild] {
			maxPos = leftChild
		}

		if rightChild <= n && slice[maxPos] < slice[rightChild] {
			maxPos = rightChild
		}

		if maxPos == i {
			return
		} else {
			slice[maxPos], slice[i] = slice[i], slice[maxPos]
			i = maxPos
		}
	}

	return
}

func heapify1(slice []int, n, i int) {
	var maxPos = i

	for ;  ; {
		if 2*i + 1 <=n && slice[i] < slice[2*i + 1] {
			maxPos = 2 * i + 1
		}

		if 2*i + 2 <= n && slice[maxPos] < slice[2*i+2] {
			maxPos = 2 * i + 2
		}

		if maxPos == i {
			return
		} else {
			slice[maxPos], slice[i] = slice[i], slice[maxPos]
			i = maxPos
		}
	}

	return
}

func sort(slice[]int) {

	n := len(slice) - 1

	fmt.Println("Before heap sort : ", slice)

	// 1. 建堆，大顶堆，最后元素是第一个，第一个与最后一个互换，最大值就到了最后一个元素，
	// 2. 在对n-1个元素堆化，重复步骤1，当slice只有一个元素，slice即是有序的。
	BuildHeap(slice)

	for ; n>=1; n--{
		slice[1], slice[n] = slice[n], slice[1]
		heapify(slice, n-1, 1)
	}

	fmt.Println("After heap sort : ", slice)
	fmt.Println()

	return
}

func sort1(slice[]int) {

	n := len(slice) - 1

	fmt.Println("Before heap sort : ", slice)

	// 1. 建堆，大顶堆，最后元素是第一个，第一个与最后一个互换，最大值就到了最后一个元素，
	// 2. 在对n-1个元素堆化，重复步骤1，当slice只有一个元素，slice即是有序的。
	BuildHeap1(slice)

	for ; n>=0; n--{
		slice[0], slice[n] = slice[n], slice[0]
		heapify1(slice, n-1, 0)
	}

	fmt.Println("After heap sort : ", slice)
	fmt.Println()

	return
}

// 从最后一个非叶子节点开始堆化
func BuildHeap(slice []int) {
	n := len(slice) - 1

	// 从n/2 + 1 到n都是叶子节点
	// 从2/n到根节点，挨个堆化即可。
	for i:=n/2;i >= 1;i-- {
		heapify(slice, n, i)
	}

	return
}


// 从最后一个非叶子节点开始堆化，数组下标为0
func BuildHeap1(slice []int) {
	n := len(slice) - 1

	// 从n/2 + 1 到n都是叶子节点
	// 从2/n到根节点，挨个堆化即可。
	for i:=n-1/2;i >= 0;i-- {
		heapify1(slice, n, i)
	}

	return
}


func main () {

	heap := new(Heap)
	heap.cap = 11
	heap.slice = make([]int, 10)
	heap.count = 0

	heap.GenerateData()
	heap.Printf()

	BuildHeap([]int{0, 4, 11, 8, 9, 7, 5})

	// 数组下表从0开始
	for i:= 8; i<=28; i++ {
		slice := util.GenerateData(i)
		sort1(slice)
		util.CheckResult(slice)
	}

	// 数组下表从1开始
	for i:= 8; i<=28; i++ {
		slice := util.GenerateData(i)
		sort(slice)
		util.CheckResult(slice)
	}

	/*
	count := heap.count
	for i:=1; i<=count; i++ {
		heap.RemoveMaxElement()
		heap.Printf()
	}
	*/


	fmt.Println("hello world")
	return
}
