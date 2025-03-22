package rpsmeter

import (
	"sync/atomic"
)

const n = 15

type Meter struct {
	buckets [n]atomic.Int64

	lastRecorded atomic.Uint64
}

func (r *Meter) Record() {
	now := currentTimestamp.Load()
	r.resetOutdated(now)

	measureId := now % n
	r.buckets[measureId].Add(1)

	resetId := (now + 2) % n
	r.buckets[resetId].Store(0)
}

// Result returns the number of requests per second for the last 10 seconds.
func (r *Meter) Result() [10]int64 {
	now := currentTimestamp.Load()
	r.resetOutdated(now)

	res := [10]int64{}
	for i := 0; i < 10; i++ {
		res[i] = r.buckets[(now-2-uint64(i))%n].Load()
	}

	return res
}

func (r *Meter) resetOutdated(now uint64) {
	last := r.lastRecorded.Load()
	if last == 0 || last > now || now-last < 3 {
		return
	}

	if !r.lastRecorded.CompareAndSwap(last, now) {
		return
	}

	diff := int(now - 2 - last)
	if diff > 10 {
		diff = 10
	}

	for i := 0; i < diff; i++ {
		r.buckets[(now-2-uint64(i))%n].Store(0)
	}
}
