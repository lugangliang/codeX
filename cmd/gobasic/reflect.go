package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name string
	Life int
}


func (* Bird) Fly() {
	fmt.Printf("I am flying....\n")
}


func main() {
	sparrow := &Bird{
		Name: "sparrow",
		Life: 3,
	}

	sparrow.Fly()

	haio := new(Bird)
	haio.Name = "haio"
	haio.Life = 4

	s := reflect.TypeOf(sparrow).Elem()
	typeOfT := s.Name()
	for i:=0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT, f.Type.Name(), f.Offset)
	}

	s = reflect.TypeOf(haio).Elem()
	typeOfT = s.Name()
	for i:=0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT, f.Type.Name(), f.Tag)
	}

	return
}