package main

import "runtime"

type Person struct {
	phone *Phone
}

type Phone struct {
	money int
}

func main() {
	chao := new(Person)
	iphone := &Phone{money: 6599}
	chao.phone = iphone
	huawei := &Phone{money: 5899}
	chao.phone = huawei

	runtime.GC()
}
