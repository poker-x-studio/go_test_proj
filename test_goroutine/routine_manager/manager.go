/*
功能：goroutine 管理
说明:
*/
package routine_manager

import (
	"sync"
)

type GoHandler func()

type RoutineManager struct {
	go_handlers []GoHandler
	wg          sync.WaitGroup
	exit_ch     chan struct{}
}

var instance *RoutineManager

// 单例
func Instance() *RoutineManager {
	if instance == nil {
		instance = &RoutineManager{}
		instance.init()
	}
	return instance
}

// 初始化
func (r *RoutineManager) init() {
	r.go_handlers = make([]GoHandler, 0)
	r.exit_ch = make(chan struct{}, 0)
}

func (r *RoutineManager) Go(f func()) {
	r.go_handlers = append(r.go_handlers, f)
	r.wg.Add(1)
	go f()
}

func (r *RoutineManager) Go_external(f func()) {
	r.go_handlers = append(r.go_handlers, f)
	r.wg.Add(1)
	go func() {
		f()
	}()
}

func (r *RoutineManager) Go_exit() {
	r.wg.Done()
}

// 通知结束所有goroutine
func (r *RoutineManager) Notify_exit() {
	len := len(r.go_handlers)
	r.exit_ch = make(chan struct{}, len)
	for i := 0; i < len; i++ {
		r.exit_ch <- struct{}{}
	}
	r.wg.Wait()
	close(r.exit_ch)
}

// 结束的通道
func (r *RoutineManager) Get_exit_ch() chan struct{} {
	return r.exit_ch
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
