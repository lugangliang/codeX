package main

import "fmt"

type Base struct {
	Name string
}

type IFoo interface {
	foo()
	hello()
}

func (*Base) foo() {
	fmt.Println("Base foo function")
}

func (*Base) bar() {
	fmt.Println("Base bar function")
}


type Foo struct {
	Base
	age string
}


func (*Foo) hello() {
	fmt.Println("In foo hello function")
}

func (*Foo) foo() {
	fmt.Println("Foo foo function")
}

func interfaceFunction() {
	var foo1 IFoo = new(Foo)
	foo1.foo()
	foo1.hello()

	foo3 := &Foo{
		age:  "18",
	}

	var ifoox IFoo = foo3
	ifoox.hello()

	return
}

func main() {

	// 多种调用方法，还有调用父类的方法
	foo := new(Foo)
	foo.foo()
	foo.bar()
	foo.Base.foo()
	foo.Base.bar()

	interfaceFunction()

	fmt.Println("Hello world.")
	return
}
