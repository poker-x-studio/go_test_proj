package main

import (
	"fmt"
	"helloworld/routine_manager"
)

func go1() {
	defer func() {
		fmt.Println("go1(),exit")
		routine_manager.Instance().Go_exit()
	}()

	for {
		select {
		case <-routine_manager.Instance().Get_exit_ch():
			fmt.Println("go1(),通知")
			return
			//default:
		}
	}
}

func go2(number int) {
	defer func() {
		fmt.Println("go2(),exit")
		routine_manager.Instance().Go_exit()
	}()

	fmt.Println(number)

	for {
		select {
		case <-routine_manager.Instance().Get_exit_ch():
			fmt.Println("go2(),通知")
			return
			//default:
		}
	}
}

func go3() {
	for {
	}
}

func main() {
	manager := routine_manager.Instance()

	manager.Go(go1)
	manager.Go(func() { go2(1) })
	manager.Go_external(go3)

	manager.Notify_exit()
	fmt.Println("main(),exit")
}
