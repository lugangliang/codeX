package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Var int
}

func MergeSortList(list1, list2 *ListNode) (head *ListNode){
	head = &ListNode{
		Var:  0,
	}

	result := head

	for ;list1 != nil && list2 != nil ; {
		if list1.Var <= list2.Var {
			head.Next = list1
			head = head.Next
			list1 = list1.Next
		} else {
			head.Next = list2
			head = head.Next
			list2 = list2.Next
		}
	}

	if list1 != nil && list2 == nil {
		head.Next = list1
	} else if list1 == nil && list2 != nil {
		head.Next = list2
	}

	return result.Next
}

func InitList(array []int) (list *ListNode){
	first := true
	var tail *ListNode
	for _, value := range array {
		temp := &ListNode{
			Var: value,
		}

		if first == true {
			list = temp
			tail = list
			first = false
			continue
		}

		if tail != nil {
			tail.Next = temp
			tail = temp
		}
	}

	return list
}

func ListPrint(head *ListNode) {
	if head == nil {
		return
	}

	for ;head != nil; {
		fmt.Println(head.Var)
		head = head.Next
	}

	return
}

func main() {
	listHead1 := InitList([]int{3,4,7,10,30})
	listHead2 := InitList([]int{1,4,10,12,29})

	//ListPrint(listHead1)
	//ListPrint(listHead2)

	ListPrint(MergeSortList(listHead1, listHead2))

	fmt.Println("Hello world")
	return
}
