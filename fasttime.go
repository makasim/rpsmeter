package rpsmeter

import (
	"sync"
	"sync/atomic"
	"time"
)

var initOnce sync.Once
var currentTimestamp *atomic.Uint64

// inspired by VictoriaMetrics lib/fasttime
// https://github.com/VictoriaMetrics/VictoriaMetrics/blob/e950846534ea594940a0f69845ef60a053c98ffd/lib/fasttime/fasttime.go#L8
func Init() {
	initOnce.Do(func() {
		go func() {
			for {
				if ms := time.Now().UnixMilli() % 1000; ms > 450 && ms < 550 {
					break
				}

				time.Sleep(time.Millisecond * 80)
			}

			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			for tm := range ticker.C {
				t := uint64(tm.Unix())
				currentTimestamp.Store(t)
			}
		}()

		currentTimestamp = func() *atomic.Uint64 {
			var x atomic.Uint64
			x.Store(uint64(time.Now().Unix()))
			return &x
		}()
	})
}
