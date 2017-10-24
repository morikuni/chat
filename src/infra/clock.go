package infra

import (
	"time"
)

func NewClock() Clock {
	return clock{}
}

type Clock interface {
	Now() time.Time
}

type clock struct{}

func (c clock) Now() time.Time {
	return time.Now()
}
