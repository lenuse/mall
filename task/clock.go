package task

import (
	"fmt"
	"math"
	"time"

	"github.com/ryszard/goskiplist/skiplist"
)

const UNTOUCHED = time.Duration(math.MaxInt64)


type Clock struct {
	seq       int                //序列
	taskQueue *skiplist.SkipList //跳表
	total     uint64             //总任务数
}

func NewClock() *Clock {
	clock := &Clock{
		taskQueue: skiplist.NewIntMap(),
	}
	clock.start()
	return clock
}

//开启时间
func (c *Clock) start() {
	emptyTask := task{
		createdAt:    time.Now(),
		intervalTime: UNTOUCHED,
		event: func() {
			fmt.Println("start clock")
		},
	}
	_, inserted := c.AddTask(emptyTask.createdAt, emptyTask.intervalTime, emptyTask.event)
	if !inserted {
		panic("[clock] internal error.Reason cannot insert job.")
	}
	go c.schedule()

}

func (c *Clock) AddTask(createdAt time.Time, actionInterval time.Duration, event func()) (t *task, inserted bool) {
	c.seq ++
	t = &task{
		id:           c.seq,
		createdAt:    createdAt,
		intervalTime: actionInterval,
		message:      make(chan Tasks, 10),
		event:        event,
		clock:        c,
	}
	c.taskQueue.Set(t.id, t)
	inserted = true
	return
}

func (c *Clock) schedule() {
	var (
		timeout time.Duration
		t     *task
		timer = newSafeTimer(UNTOUCHED)
	)
	i := c.taskQueue.Iterator()
	defer timer.Stop()
	for {
		if  ok := i.Next(); ok {
			t = i.Value().(*task)
			timeout = t.intervalTime
			timer.Reset(timeout)
		}
		select {
		case <-timer.C:
			t.run(true)
		default:

		}
	}
}
