package rpsmeter

import (
	"reflect"
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

func TestMeterRPS(t *testing.T) {
	f := func(rps float64, expectedTotal int64, expectedBreakdown [10]int64) {
		t.Helper()

		synctest.Run(func() {

			stopCh := make(chan struct{})
			defer close(stopCh)

			var m Meter
			go func() {
				startAtSameMS()

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

			actualTotal, actualBreakdown := m.Result()
			if actualTotal != expectedTotal {
				t.Fatalf("unexpected meter result total; got %v; want %v", actualTotal, expectedTotal)
			}
			if !reflect.DeepEqual(actualBreakdown, expectedBreakdown) {
				t.Fatalf("unexpected meter result rps breakdown; got %v; want %v", actualBreakdown, expectedBreakdown)
			}
		})
	}

	f(20, 390, [10]int64{20, 20, 20, 20, 20, 20, 20, 20, 20, 20})
	f(10, 195, [10]int64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10})
	f(1, 19, [10]int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	f(0.8, 15, [10]int64{0, 1, 1, 1, 1, 0, 1, 1, 1, 1})
	f(0.66666, 13, [10]int64{1, 0, 2, 1, 0, 1, 1, 0, 1, 1})
	f(0.5, 9, [10]int64{1, 1, 1, 0, 1, 0, 1, 0, 1, 0})
	f(0.33333, 6, [10]int64{1, 0, 0, 1, 0, 0, 1, 0, 0, 1})
	f(0.2, 3, [10]int64{0, 0, 0, 1, 0, 0, 0, 0, 1, 0})
	f(0.1, 1, [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 1, 0})
}

func TestMeterResetOutdated(t *testing.T) {
	f := func(wait time.Duration, expectedBreakdown [10]int64) {
		t.Helper()

		synctest.Run(func() {
			var m Meter

			stopCh := make(chan struct{})
			go func() {
				startAtSameMS()

				genT := time.NewTicker(time.Millisecond * 100)
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
			time.Sleep(time.Second * 15)
			close(stopCh)
			time.Sleep(wait)

			_, actualBreakdown := m.Result()
			if !reflect.DeepEqual(actualBreakdown, expectedBreakdown) {
				t.Fatalf("unexpected meter result rps breakdown; got %v; want %v", actualBreakdown, expectedBreakdown)
			}
		})
	}

	f(0, [10]int64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10})
	f(time.Second, [10]int64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10})
	f(time.Second*2, [10]int64{0, 10, 10, 10, 10, 10, 10, 10, 10, 10})
	f(time.Second*3, [10]int64{0, 0, 10, 10, 10, 10, 10, 10, 10, 10})
	f(time.Second*4, [10]int64{0, 0, 0, 10, 10, 10, 10, 10, 10, 10})
	f(time.Second*5, [10]int64{0, 0, 0, 0, 10, 10, 10, 10, 10, 10})
	f(time.Second*6, [10]int64{0, 0, 0, 0, 0, 10, 10, 10, 10, 10})
	f(time.Second*7, [10]int64{0, 0, 0, 0, 0, 0, 10, 10, 10, 10})
	f(time.Second*8, [10]int64{0, 0, 0, 0, 0, 0, 0, 10, 10, 10})
	f(time.Second*9, [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 10, 10})
	f(time.Second*10, [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 10})
	f(time.Second*11, [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*12, [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*13, [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*14, [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*15, [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func TestMeterResetOutdatedAndRecord(t *testing.T) {
	f := func(wait time.Duration, expectedBreakdown [10]int64) {
		t.Helper()

		synctest.Run(func() {
			var m Meter

			stopCh := make(chan struct{})
			go func() {
				startAtSameMS()

				genT := time.NewTicker(time.Millisecond * 100)
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
			time.Sleep(time.Second * 15)
			close(stopCh)
			time.Sleep(wait)
			m.Record()
			time.Sleep(time.Second * 2)

			_, actualBreakdown := m.Result()
			if !reflect.DeepEqual(actualBreakdown, expectedBreakdown) {
				t.Fatalf("unexpected meter result rps breakdown; got %v; want %v", actualBreakdown, expectedBreakdown)
			}
		})
	}

	f(0, [10]int64{1, 10, 10, 10, 10, 10, 10, 10, 10, 10})
	f(time.Second, [10]int64{1, 0, 10, 10, 10, 10, 10, 10, 10, 10})
	f(time.Second*2, [10]int64{1, 0, 0, 10, 10, 10, 10, 10, 10, 10})
	f(time.Second*3, [10]int64{1, 0, 0, 0, 10, 10, 10, 10, 10, 10})
	f(time.Second*4, [10]int64{1, 0, 0, 0, 0, 10, 10, 10, 10, 10})
	f(time.Second*5, [10]int64{1, 0, 0, 0, 0, 0, 10, 10, 10, 10})
	f(time.Second*6, [10]int64{1, 0, 0, 0, 0, 0, 0, 10, 10, 10})
	f(time.Second*7, [10]int64{1, 0, 0, 0, 0, 0, 0, 0, 10, 10})
	f(time.Second*8, [10]int64{1, 0, 0, 0, 0, 0, 0, 0, 0, 10})
	f(time.Second*9, [10]int64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*10, [10]int64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*11, [10]int64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*12, [10]int64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*13, [10]int64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*14, [10]int64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	f(time.Second*15, [10]int64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0})
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

func startAtSameMS() {
	for {
		if ms := time.Now().UnixMilli() % 1000; ms > 450 && ms < 550 {
			break
		}

		time.Sleep(time.Millisecond * 80)
	}
}
