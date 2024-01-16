package format

import (
	"fmt"
	"time"
)

type timer struct {
	start time.Time
}

func TimerStart() *timer {
	t := timer{}
	t.start = time.Now()
	return &t
}

func (_t *timer) End() {
	end := time.Now()
	InfoMessage("Cost ", fmt.Sprint(end.Sub(_t.start)))
}

func UnixTime() int64 {
	return time.Now().Unix()
}

func Timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
