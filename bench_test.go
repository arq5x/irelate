package irelate

import (
	"github.com/brentp/ififo"
	"testing"
)

func benchmarkStreams(nStreams int, b *testing.B) {

	for n := 0; n < b.N; n++ {
		streams := make([]RelatableChannel, 0)
		f := "data/test.bed.gz"
		s := ififo.NewIFifo(3000, func() interface{} { return &Interval{} })

		for i := 0; i < nStreams; i++ {
			streams = append(streams, Streamer(f, s))
		}

		merged := Merge(streams...)

		for interval := range IRelate(merged, CheckRelatedByOverlap, false, 0, s) {
			s.Put(interval)
		}

	}
}

func Benchmark2Streams(b *testing.B) { benchmarkStreams(2, b) }
func Benchmark3Streams(b *testing.B) { benchmarkStreams(3, b) }