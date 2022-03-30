package main

import (
	"fmt"
	"strconv"
)

type ArrayError int

func (ae ArrayError) Error() string {
	return "error" + strconv.FormatInt(int64(ae), 10)
}

type DyArray struct {
	Data []interface{}
	len int
	cap int
}


func (da *DyArray) resize() {
	newArray := make([]interface{}, 2*da.cap)
	for index, value := range da.Data{
		newArray[index] = value
	}
	da.Data = newArray
	da.cap = 2*da.cap
}


func (da *DyArray) AddItem (pos, value interface{}) {
	if da.len == da.cap{
		da.resize()
	}

	da.Data[da.len] = value
	da.len = da.len + 1

	return
}

func (da *DyArray) DeleteItemByIndex(index int) (err error)  {
	if da.isIndexOutOfBound(index) {
		return ArrayError(index)
	}

	da.Data[index-1] = 0

	return nil
}

func (da *DyArray) DataLen() int {
	return da.len
}

func (da *DyArray) isIndexOutOfBound(index int) bool{
	if index >= da.len {
		return true
	}

	return false
}

func MergeTwoSortedArray(a1 []int, a2 []int) []int {
	lenA1 := len(a1)
	lenA2 := len(a2)

	newArray := make([]int, lenA2+lenA1)

	i,j,k:=0,0,0
	for ; i<lenA1&&j<lenA2;k++{

		// 相同数值以a2为大
		if a1[i] <= a2[j] {
			newArray[k] = a1[i]
			// i指向下一个元素
			i++
		} else {
			newArray[k] = a2[j]
			// j指向下一个元素
			j++
		}
	}

	// 循环退出有几种情况，i == lenA1 || j == lenA2 || i == lenA1 && j == lenA2
	if i < lenA1 {
		copy(newArray[k:], a1[i:])
	}

	if j < lenA2{
		copy(newArray[k:], a2[j:])
	}

	return newArray
}

func main() {
	dyArray := &DyArray{
		Data: make([]interface{}, 10),
		len:   0,
		cap:   10,
	}

	// 这样直接赋值会导致爆数组越界错误
	/*
	for i:=0; i<20;i++ {
		dyArray.array[i] = i
		fmt.Println(dyArray.array, len(dyArray.array), cap(dyArray.array))

	}
	 */

	// 使用append才会触发slice自动扩容
	for i:=0; i<6;i++ {
		dyArray.AddItem(i, i)
	}

	fmt.Println(dyArray.Data, len(dyArray.Data), cap(dyArray.Data))




	return
}