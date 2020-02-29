package task

import (
	"log"
	"runtime/debug"
	"sync/atomic"
	"time"
)

type Tasks interface {
	Get() <-chan Tasks //Get a Chan，which can get message if tasks is executed
	Cancel()           // cancel task
	isAvailable() bool     // task is Available
}

type task struct {
	id           int        //任务id
	intervalTime time.Duration //间隙时间
	event        func()        //事件方法
	message      chan Tasks    //消息
	createdAt    time.Time     //创建时间
	state        int32
	clock		*Clock
}

func (t *task) Get() <-chan Tasks {
	return t.message
}

//func (t task) LessThan(other skiplist.Ordered) bool {
//	return t.id < other.(task).id
//}

func (t *task) isAvailable() bool {
	return t.state == 1
}

func (t *task) Cancel()  {
	if atomic.CompareAndSwapInt32(&t.state, 1, 0) {
		t.clock = nil
	}
}

func (t *task) run(async bool)  {
	if async {
		go safeCallback(t.event)
	} else {
		safeCallback(t.event)
	}
	select {
	case t.message <- t:
	default:
		
	}
}

func safeCallback(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[clock] recovering reason is %+v. More detail:", err)
			log.Println(string(debug.Stack()))
		}
	}()
	fn()
}
