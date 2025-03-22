package rpsmeter_test

import (
	"testing"

	"github.com/makasim/rpsmeter"
)

func BenchmarkMeter_Record(b *testing.B) {
	rpsmeter.Init()

	var m rpsmeter.Meter

	b.ReportAllocs()
	for b.Loop() {
		m.Record()
	}
}

func BenchmarkMeter_Parallel_Record(b *testing.B) {
	rpsmeter.Init()

	var m rpsmeter.Meter

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Record()
		}
	})
}

func BenchmarkMeter_Result(b *testing.B) {
	rpsmeter.Init()

	var m rpsmeter.Meter

	b.ReportAllocs()
	for b.Loop() {
		m.Result()
	}
}

func BenchmarkMeter_Parallel_Result(b *testing.B) {
	rpsmeter.Init()

	var m rpsmeter.Meter

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Result()
		}
	})
}
