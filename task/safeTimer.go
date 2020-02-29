package task

import "time"

/*
This code is based on the following resources:
source code:	https://play.golang.org/p/Ys9qqanqmU
discuss:	https://groups.google.com/forum/#!msg/golang-dev/c9UUfASVPoU/tlbK2BpFEwAJ
 */
type safeTimer struct {
	*time.Timer
	scr bool
}

//saw channel read, must be called after receiving value from safetimer chan
func (t *safeTimer) SCR() {
	t.scr = true
}

func (t *safeTimer) SafeReset(d time.Duration) bool {
	ret := t.Stop()
	if !ret && !t.scr {
		<-t.C
	}
	t.Timer.Reset(d)
	t.scr = false
	return ret
}
func newSafeTimer(d time.Duration) *safeTimer {
	return &safeTimer{
		Timer: time.NewTimer(d),
	}
}