package main

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 第23题，合并k个有序链表

import (
	"fmt"
)

var (
	max *ListNode
)

type ListNode struct {
      Val int
      Next *ListNode
 }


 func (list *ListNode) PrintList(){
 	var tmp = list
 	for ;tmp!=nil; tmp=tmp.Next{
 		fmt.Printf("%d \t ", tmp.Val)
	}
	fmt.Println()
 }

 func (list *ListNode) InsertBeforeNode(node *ListNode) *ListNode {

 	var head *ListNode
 	head = list

 	node.Next = head.Next
	head.Next = node

	return  head
 }

func insertNodeToList(value int, sentinelNode *ListNode) {

	newNode := &ListNode{Val:value}

	var tmpNode = sentinelNode.Next
	if tmpNode == nil {
		sentinelNode.Next = newNode
		return
	}

	prev := sentinelNode

	for ;tmpNode!=nil; {
		if newNode.Val <=  tmpNode.Val {
			newNode.Next = tmpNode
			prev.Next = newNode
			return
		} else {
			prev = tmpNode
			tmpNode=tmpNode.Next
		}
	}
	newNode.Next = nil
	max = newNode
	prev.Next = newNode

}

func mergeTwoSortedList(sentinelNode *ListNode, list *ListNode) {

	var tmpNode *ListNode

	for tmpNode = list ;tmpNode != nil ; tmpNode=tmpNode.Next{
		insertNodeToList(tmpNode.Val, sentinelNode)
	}

}

func mergeKLists(lists []*ListNode) *ListNode {

	sentinelNode := &ListNode{}

	for _, list := range lists {
		if list == nil {
			continue
		}
		// 第一个非空链表,哨兵节点指向第一个
		if sentinelNode.Next == nil {
			sentinelNode.Next = list
		} else {
			mergeTwoSortedList(sentinelNode, list)
		}
	}

	return sentinelNode.Next

}

// [[1,4,5],[1,3,4],[2,6]]
func initData() []*ListNode {

	listArray := make([]*ListNode,3 )
	list1 := &ListNode{
		Val: 1,
	}
	list1.InsertBeforeNode(&ListNode{Val:5})
	list1.InsertBeforeNode(&ListNode{Val:4})

	list2 := &ListNode{
		Val: 1,
	}
	list2.InsertBeforeNode(&ListNode{Val:4})
	list2.InsertBeforeNode(&ListNode{Val:3})

	list3 := &ListNode{
		Val: 2,
	}
	list3.InsertBeforeNode(&ListNode{Val:6})

	listArray[0] = list1
	listArray[1] = list2
	listArray[2] = list3

	return listArray
}

func main() {


	listArray  := initData()
	for _, list := range listArray {
		list.PrintList()
	}

	fmt.Println("after merge!")

	mergeKLists(listArray).PrintList()

	return
}