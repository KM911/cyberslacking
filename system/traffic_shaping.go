package system

import (
	"sync"
	"time"
)

func NewDebounce(_interval time.Duration) func(f func()) {
	var l sync.Mutex
	var timer *time.Timer

	return func(f func()) {
		l.Lock()
		defer l.Unlock()
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(_interval, f)
	}
}

func NewThrottle(_interval time.Duration) func(f func()) {
	var l sync.Mutex
	var timer *time.Timer

	return func(f func()) {
		l.Lock()
		defer l.Unlock()
		if timer == nil {
			timer = time.AfterFunc(_interval, func() {
				f()
				timer = nil
			})
		}
	}
}
