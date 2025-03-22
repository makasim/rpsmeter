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
	r.cleanupOutdated(now)

	stored := r.lastRecorded.Swap(now)
	if now > stored && now-stored > 2 {
		for i := stored + 1; i < now; i++ {
			r.buckets[i%n].Store(0)
		}
	}

	measureId := now % n
	r.buckets[measureId].Add(1)

	resetId := (now + 2) % n
	r.buckets[resetId].Store(0)
}

func (r *Meter) LastTenSeconds() []int64 {
	now := currentTimestamp.Load()
	r.cleanupOutdated(now)

	res := make([]int64, 10)
	for i := 0; i < 10; i++ {
		res[i] = r.buckets[(now-2-uint64(i))%n].Load()
	}

	return res
}

func (r *Meter) cleanupOutdated(now uint64) {
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

	dbg := make([]int, 0, diff)
	for i := 0; i < diff; i++ {
		dbg = append(dbg, int((now-2-uint64(i))%n))
		r.buckets[(now-2-uint64(i))%n].Store(0)
	}

	dbg = dbg[:0]
}
