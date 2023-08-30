package main

import (
	"fmt"
	"reflect"
)

type Data1 struct {
	index int32
	name  string
}

type Data2 struct {
	index int64
	name  string
}

func Print[T Data1 | Data2](d T) {
	if reflect.TypeOf(d).Name() == "Data1" {
		fmt.Println("1")
	}
}

func main() {
	//var x interface{}
	var data1 Data1
	var data2 Data2

	fmt.Println(reflect.TypeOf(data1).Kind())
	fmt.Println(reflect.TypeOf(data1).Name())
	fmt.Println(reflect.TypeOf(data1).String())
	Print(data1)

	fmt.Println(reflect.TypeOf(data2).Kind())
	fmt.Println(reflect.TypeOf(data2).Name())
	fmt.Println(reflect.TypeOf(data2).String())
	Print(data2)

	if reflect.TypeOf(data1).String() == "main.Data1" {
		fmt.Println(" y")
	}

}
