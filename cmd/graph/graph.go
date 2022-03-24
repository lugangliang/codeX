package main

import (
	"container/list"
	"fmt"
)

var (
	found int
	queue *list.List
	v int
)

func InitData() []*list.List {
	v = 10

	adj := make([] *list.List, v)

	for i:=0; i<v; i++ {
		adj[i] = list.New()
	}

	insert(4,3, adj)
	insert(4,8, adj)
	insert(4,9, adj)
	insert(3,6, adj)
	insert(3,9, adj)
	insert(6,7, adj)
	insert(9,7, adj)
	insert(9,2, adj)
	insert(8,1, adj)
	insert(8,9, adj)
	insert(7,5, adj)
	insert(2,0, adj)
	insert(2,5, adj)
	insert(5,0, adj)

	queue = list.New()

	/*
	queue.PushBack("h")
	queue.PushBack("e")
	queue.PushBack("l")
	queue.PushBack("l")
	queue.PushBack("o")


	for ;; {
		e:=queue.Front()
		if e!=nil {
			fmt.Println(queue.Remove(e))
		} else {
			break
		}
	}
*/

	return adj
}

func insert(i,j  int, adj []*list.List){
	adj[i].PushBack(j)
	adj[j].PushBack(i)
}

func PrintAdj(adj[]*list.List) {
	for index, l := range adj{
		fmt.Printf("%d : ", index)
		for e:=l.Front(); e!=nil; e=e.Next(){
			fmt.Printf("%d \t", e.Value.(int))

		}
		fmt.Println()
	}
}

func PrintPrev(prev []int, start int) {
	if start == -1 || prev[start] == -1{
		fmt.Println(start)
		return
	}


	PrintPrev(prev, prev[start])
	fmt.Println(start)
	return
}

func Bfs(start, end int, adj []*list.List) {

	queue.PushBack(start)

	visted := make([]bool, v)
	visted[start] = true

	prev := make([]int, v)
	for index,_ := range  prev{
		prev[index] = -1
	}

	for {
		e := queue.Front()
		if e != nil {
			value, ok := e.Value.(int)
			if !ok {
				return
			}
			// 遍历邻接表
			for edge := adj[value].Front(); edge != nil; edge = edge.Next() {
				edgeValue, ok := edge.Value.(int)
				if !ok {
					return
				}

				if !visted[edgeValue] {
					prev[edgeValue] = value
					if edgeValue != end {
						visted[edgeValue] = true
						queue.PushBack(edgeValue)
					} else {
						PrintPrev(prev, end)
						return
					}
				}
			}
			queue.Remove(e)
		} else {
			return
		}

	}

}

func main() {
	adj := InitData()

	PrintAdj(adj)

	Bfs(0,8, adj)

	fmt.Println("hello world", found)

	return
}