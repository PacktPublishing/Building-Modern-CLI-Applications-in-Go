package examples

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

type runner struct {
	Run     func() bool
	limiter *rate.Limiter
}

func Limit() {
	thing := runner{}
	start := time.Now()
	thing.Run = func() bool {
		if thing.limiter.Allow() {
			fmt.Println(time.Now())
			return false
		}
		if time.Since(start) > 30*time.Second {
			return true
		}
		return false
	}
	thing.limiter = rate.NewLimiter(forEvery(1, 5*time.Second), 1)
	for {
		if thing.Run() {
			break
		}
	}
}

func forEvery(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}
