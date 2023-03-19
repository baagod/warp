package beam

import (
	"time"
)

type Timer struct {
	timer *time.Timer
	f     func()
	done  chan bool
}

// After 返回一个定时器，f 将等待 d 结束后在协程中执行。
// 调用 Stop 可以停止定时器并中止协程。
func After(d time.Duration, f func()) *Timer {
	t := &Timer{timer: time.NewTimer(d)}
	t.f = f
	t.done = make(chan bool, 1)

	goFunc(t)
	return t
}

// Stop 如果调用停止计时器，则返回 true，如果计时器已过期或已停止，则返回 false。
func (t *Timer) Stop() (stopped bool) {
	// t.timer.Stop 内部不会关闭 t.timer.C 通道
	if stopped = t.timer.Stop(); stopped {
		t.done <- true // 退出协程
	}
	return
}

// Reset 将定时器更改为在持续时间 d 后过期。
func (t *Timer) Reset(d time.Duration) bool {
	if !t.timer.Stop() { // 如果计时器过期或已停止
		goFunc(t) // 重新启动协程方法
	}
	return t.timer.Reset(d) // 重置通道 t.timer.C 的发送时间为 d
}

func goFunc(t *Timer) {
	go func() {
		select {
		case <-t.done: // 释放协程
		case <-t.timer.C:
			if t.f != nil {
				t.f()
			}
		}
	}()
}
