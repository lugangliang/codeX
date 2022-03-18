package main

import (
	"fmt"
)

// 批量初始化
var (
	e = 10
	f = 11
)

type PersonInfo struct {
	ID string
	Name string
	Address string
}

const (
	size int64 = 1024
)

// 变量定义，其实这里应该要简化的，需要记住三种。
func paramDefine() {
	var a int = 3
	var b = 4
	c := 5
	// 先初始化，再赋值
	var d int

	a = d
	b = a
	b = c
	a = b
	return
}

func strFunction() {
	x := "hello"
	y := "world"
	// 字符串拼接
	s := x + y
	fmt.Printf("s is %s, len is %d, first is %d\n", s, len(s), s[0])

	// 字符串遍历
	for i:=0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	// golang的range实现
	for key, value := range s {
		fmt.Println(key, value)
	}
}

func arrayFuntion(array []int) {
	a := [3]int{1,2,3}
	a[1] = 3
	b := [3][4]int{{1,2},{2,3}}
	b[1][2] = 5
	fmt.Println("in arrayFunction length is", len(array))

	return
}

func sliceFunction() {
	// 数组是值传递，切片是指针传递
	// 写死容量，就是数组，不写死，就是切片

	// 1.基于数组创建切片
	var myArray [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var mySlice []int = myArray[:5]
	for _, value := range mySlice{
		fmt.Println(value)
	}
	mySlice = myArray[:]
	mySlice = myArray[2:]
	mySlice = myArray[:4]

	// 2. 使用make创建切片
	var makeSlice = make([]int, 4, 20)

	makeSlice = append(makeSlice, 1, 2, 3, 4)
	// 直接追加切片，后面需要再加...
	makeSlice = append(makeSlice, []int{10, 11, 12}...)
	mySlice2 := []int{22, 23, 24}
	makeSlice = append(makeSlice, mySlice2...)
	// 超出20容量后，自动扩容
	makeSlice = append(makeSlice, []int{2,3,4,33,44,5,5,6,6,7,7,8,89,9,9,0,3,3,3}...)

	for i,j := range makeSlice{
		fmt.Println(i, j)
	}

	// 两个内置函数len和cap获取切片的信息
	fmt.Println("len(makeSlice) :", len(makeSlice))
	fmt.Println("cap(makeSlice) :", cap(makeSlice))

	// 内置函数copy,只赋值最少的那个slice
	slice1 := []int{1,2,3,4,5}
	slice2 := []int{5,4,3}
	copy(slice1, slice2)
	copy(slice2, slice1)
	fmt.Println(slice2, slice1)

	return
}

func mapFunction() {
	// 创建map
	myMap := make(map[string]PersonInfo, 100)
	myMap["lugl"] = PersonInfo{"12", "lugangliang", "zz"}
	myMap["xun"] = PersonInfo{"34", "xuning", "aa"}

	// 查找map
	persion, ok := myMap["xx"]
	if !ok {
		fmt.Println("xx is not found in map")
	} else {
		fmt.Println(persion)
	}

	// 删除map
	delete(myMap, "abc")
	delete(myMap, "xun")

	for key,value := range myMap{
		fmt.Println(key, value)
	}

	return
}

func switchFuntion(a int) {
	switch  a {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
		fallthrough
	case 3,4,5:
		fmt.Println(3,4,5)
	default:
		fmt.Println("default")
	}

	return
}

func loopFunction() {
	// 只有for

	// 无限循环
	for {
		a := 10
		a += 10
		if a > 15 {
			break
		}
	}

    // 有限循环
    var sum int
    for i:=0; i<10; i++{
    	sum +=i
	}
	fmt.Println("sum is ", sum)

	return
}

// random parameter function
// interface{}，就可以传入任意的类型了
func randParmeterFunction(args ...interface{}) {
	for _, arg := range args{
		// .(type)可以返回类型
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a  string value")
		default:
			fmt.Println(arg, "is an unknown type")
		}
	}
	return
}

// go的匿名函数与闭包的用法,适当使用闭包，代码更简洁，可读性也更好
func closureOperation() {
	// 匿名函数
	b := func(a1, a2 int) int {
		fmt.Println("ano function call ", a1 + a2)
		return a1 + a2
	}
	b(3, 5)

	// 闭包的应用，函数可以作为参数，也可以作为函数值
	var j = 5
	// 执行完后，a就是一个函数指针了
	a := func() func(op string) int {
		i := 10
		return func(op string) int {
			switch op {
			case "add":
				return i + j
			case "del":
				return j - i
			case "mut":
				return j*i
			default:
				return -1
			}
		}
	}()

	fmt.Println("add", a("add"))

	// 改变j的值，再次执行a(),得到的结果也是不同的。
	j = 8
	fmt.Println("mut", a("mut"))

	return
}

// 函数返回值可以不显式赋值
func getName() (firstName, NextName, nickName string) {
	return "mark", "lug", "gangliang"
}

func deferOperation(){
	defer func() {
		fmt.Println("In defer operation.")
	}()

	fmt.Println("in defer function..")
	return
}

// 结构体的四种初始化方法
func structOperation() {
	person1 := new(PersonInfo)
	person2 := &PersonInfo{}
	person3 := &PersonInfo{"123", "fs", "sss"}
	person4 := &PersonInfo{
		ID:      "d",
		Name:    "a",
		Address: "c",
	}
	fmt.Println(person1, person2, person3, person4)

	return
}

func main() {
	paramDefine()

	structOperation()

	strFunction()

	arrayFuntion([]int{1,2,3,4})

	sliceFunction()

	mapFunction()

	switchFuntion(1)

	loopFunction()

	randParmeterFunction(1, 3.6, "aaaa", -4)

	closureOperation()

	deferOperation()



	_, _, name := getName()
	fmt.Printf("Hello world %s %d %d %d\n", name, e, f, size)

	return
}