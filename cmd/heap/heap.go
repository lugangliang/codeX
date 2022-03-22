package main

import (
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

func main () {

	heap := new(Heap)
	heap.cap = 11
	heap.slice = make([]int, 10)
	heap.count = 0

	heap.GenerateData()
	heap.Printf()

	count := heap.count
	for i:=1; i<=count; i++ {
		heap.RemoveMaxElement()
		heap.Printf()
	}

	fmt.Println("hello world")
	return
}
