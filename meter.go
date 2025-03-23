package rpsmeter

import (
	"sync/atomic"
	"time"
)

const n = 15

type Meter struct {
	total   atomic.Int64
	buckets [n]atomic.Int64

	lastRecorded atomic.Int64
}

// Record increments the total request count and updates the current second's bucket, resetting outdated ones.
func (m *Meter) Record() {
	m.total.Add(1)

	now := time.Now().Unix()
	m.resetOutdated(now)

	measureId := now % n
	m.buckets[measureId].Add(1)

	resetId := (now + 2) % n
	m.buckets[resetId].Store(0)
}

// Result returns the total number of requests and a breakdown of requests per second over the last 10 seconds.
// The per-second request counts are provided as an array, where the most recent second is at index 0.
func (m *Meter) Result() (int64, [10]int64) {
	now := time.Now().Unix()
	m.resetOutdated(now)

	res := [10]int64{}
	for i := 0; i < 10; i++ {
		res[i] = m.buckets[(now-2-int64(i))%n].Load()
	}

	return m.total.Load(), res
}

func (m *Meter) resetOutdated(now int64) {
	last := m.lastRecorded.Load()
	if last == 0 || last > now || now-last < 3 {
		m.lastRecorded.Store(now)
		return
	}

	if !m.lastRecorded.CompareAndSwap(last, now) {
		return
	}

	diff := int(now - last)
	if diff > 15 {
		diff = 15
	}

	dbg := []int{}
	for i := 0; i < diff; i++ {
		dbg = append(dbg, int((now-int64(i))%n))
		m.buckets[(now-int64(i))%n].Store(0)
	}
}
