package rpsmeter

import (
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"
)

func TestMeterRPS(t *testing.T) {
	var x atomic.Uint64
	currentTimestamp = func() *atomic.Uint64 {
		return &x
	}()

	f := func(rps float64, expected [10]int64) {
		t.Helper()

		synctest.Run(func() {

			stopCh := make(chan struct{})
			defer close(stopCh)

			x.Store(uint64(time.Now().Unix()))
			startFastTime(stopCh)

			var m Meter
			go func() {
				genT := time.NewTicker(time.Microsecond * time.Duration(1/rps*1000000))
				defer genT.Stop()

				for {
					select {
					case <-stopCh:
						return
					case <-genT.C:
						m.Record()
					}
				}
			}()

			time.Sleep(time.Second * 20)

			actual := m.Result()
			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("unexpected meter result; got %v; want %v", actual, expected)
			}
		})
	}

	f(20, [10]int64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20})
	f(10, [10]int64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10})
	f(1, [10]int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	f(0.66666, [10]int64{2, 1, 0, 1, 1, 0, 1, 1, 0, 1})
	f(0.5, [10]int64{1, 1, 1, 0, 1, 0, 1, 0, 1, 0})
	f(0.33333, [10]int64{2, 0, 0, 1, 0, 0, 1, 0, 0, 1})
	f(0.2, [10]int64{0, 0, 0, 1, 0, 0, 0, 0, 1, 0})
	f(0.1, [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 1, 0})
}

func TestMeterResetOutdated(t *testing.T) {
	var x atomic.Uint64
	currentTimestamp = func() *atomic.Uint64 {
		return &x
	}()

	synctest.Run(func() {

		stopCh := make(chan struct{})
		defer close(stopCh)

		x.Store(uint64(time.Now().Unix()))
		startFastTime(stopCh)

		var m Meter

		m.Record()
		time.Sleep(time.Second * 10)

		expected := [10]int64{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}
		actual := m.Result()
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("unexpected meter result; got %v; want %v", actual, expected)
		}
	})
}

func TestMeter_Concurrency(t *testing.T) {
	var m Meter

	stopCh := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-stopCh:
					return
				default:
					m.Record()
					time.Sleep(time.Millisecond * 10)
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-stopCh:
					return
				default:
					m.Result()
					time.Sleep(time.Millisecond * 10)
				}
			}
		}()
	}

	time.Sleep(time.Second * 2)
	close(stopCh)

	wg.Wait()
}

func startFastTime(stopCh chan struct{}) {
	go func() {
		for {
			if ms := time.Now().UnixMilli() % 1000; ms > 450 && ms < 550 {
				break
			}

			time.Sleep(time.Millisecond * 80)
		}

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-stopCh:
				return
			case tm := <-ticker.C:
				t := uint64(tm.Unix())
				currentTimestamp.Store(t)
			}
		}
	}()
}
