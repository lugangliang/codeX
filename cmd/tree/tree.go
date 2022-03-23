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

	for _, value := range heap.slice {
		a += " " + strconv.Itoa(value)
	}

	return a
}



func (heap *Heap) Printf() {
	fmt.Println(heap.slice)
}

func (heap *Heap) Insert(value int) {
	if heap.count >= heap.cap {
		return
	}

	heap.count++
	i := heap.count
	heap.slice[i] = value

	for ;i >= 1; {
		if heap.slice[i] < heap.slice[i/2] {
			heap.slice[i] = heap.slice[i/2]
			i = i/2
		} else {
			break
		}
	}

	return
}

func (heap *Heap) GenerateData() {
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(10)
	heap.Insert(8)
	heap.Insert(6)

	return
}

func main () {

	heap := new(Heap)
	heap.cap = 11
	heap.slice = []int{}
	heap.count = 0

	heap.GenerateData()

	heap.Printf()

	fmt.Println("hello world")
	return
}
