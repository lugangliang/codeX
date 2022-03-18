package main

import (
	"fmt"
	"reflect"
)


func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}

	resultChan <- sum
	return
}

func main() {
	// 变量的初始化方法
	values := []int{1,2,3,4,5,1,2,3,4,5}

	// make的使用方法
	resultChan := make(chan int, 2)

	// goroutine和数组切片的使用方法
	go sum(values[0:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)

	// channel的使用方法
	sum1, sum2 := <-resultChan, <-resultChan

	fmt.Println("Result is ", sum1, sum2, sum1 + sum2)

	var name = "abcdABCD1234"
	for _, char := range  name {
		fmt.Printf("%v\n", reflect.TypeOf(char))
	}

	var hoby []byte = []byte{1, 97, 98}
	for _, char := range hoby {
		fmt.Printf("%v\n", char)
	}

	return
}
